package plaid

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/MicahParks/keyfunc"
	"github.com/golang-jwt/jwt/v4"
	plaidSdk "github.com/plaid/plaid-go/v10/plaid"
	"log"
	"strconv"
	"strings"
	"time"
	"wallit/ent"
	"wallit/ent/plaiditem"
	"wallit/ent/transaction"
	"wallit/ent/transactionsync"
	"wallit/internal/config"
	"wallit/internal/db"
	"wallit/internal/notification"
	"wallit/internal/spending_category"
)

type ServiceConfig struct {
	webHookUrl string
}

type Plaid struct {
	client                  *plaidSdk.APIClient
	config                  *ServiceConfig
	spendingCategoryService *spending_category.Service
	notificationService     *notification.Service
	db                      *ent.Client
}

func New(
	spendingCategoryService *spending_category.Service,
	notificationService *notification.Service,
	db *ent.Client,
) *Plaid {
	configuration := plaidSdk.NewConfiguration()
	configuration.AddDefaultHeader("PLAID-CLIENT-ID", ClientId())
	configuration.AddDefaultHeader("PLAID-SECRET", Secret())
	configuration.AddDefaultHeader("Plaid-Version", "2020-09-14")
	configuration.UseEnvironment(Environment())

	return &Plaid{
		client: plaidSdk.NewAPIClient(configuration),
		config: &ServiceConfig{
			webHookUrl: config.Get[string]("PLAID_WEBHOOK_URL"),
		},
		spendingCategoryService: spendingCategoryService,
		notificationService:     notificationService,
		db:                      db,
	}
}

func (p *Plaid) WithTransactionClient(db *ent.Client) *Plaid {
	return &Plaid{
		client:                  p.client,
		config:                  p.config,
		spendingCategoryService: p.spendingCategoryService,
		notificationService:     p.notificationService,
		db:                      db,
	}
}

func (p *Plaid) CreateLinkToken(userId int) (string, error) {
	// @TODO: in the future add customisation and improve the conversion

	requestData := plaidSdk.LinkTokenCreateRequestUser{
		ClientUserId: strconv.Itoa(userId),
	}
	request := plaidSdk.NewLinkTokenCreateRequest(
		"Wallit Platform",
		"en",
		[]plaidSdk.CountryCode{plaidSdk.COUNTRYCODE_US},
		requestData,
	)
	request.SetProducts([]plaidSdk.Products{plaidSdk.PRODUCTS_TRANSACTIONS})
	//request.SetLinkCustomizationName("default")
	request.SetWebhook(p.config.webHookUrl)

	res, _, err := p.client.PlaidApi.
		LinkTokenCreate(context.Background()).
		LinkTokenCreateRequest(*request).
		Execute()

	if err != nil {
		log.Print(err)

		return "", err
	}

	return res.GetLinkToken(), nil
}

func (p *Plaid) ExchangePublicTokenIntoAccessToken(ctx context.Context, publicToken string) (plaidSdk.ItemPublicTokenExchangeResponse, error) {
	tokenExchangeReq := plaidSdk.NewItemPublicTokenExchangeRequest(publicToken)
	tokenExchangeRes, _, err := p.client.PlaidApi.
		ItemPublicTokenExchange(ctx).
		ItemPublicTokenExchangeRequest(*tokenExchangeReq).
		Execute()

	return tokenExchangeRes, err
}

func (p *Plaid) GetItemInfo(ctx context.Context, accessToken string) (plaidSdk.ItemGetResponse, error) {
	request := plaidSdk.NewItemGetRequest(accessToken)
	res, _, err := p.client.PlaidApi.
		ItemGet(ctx).
		ItemGetRequest(*request).
		Execute()

	return res, err
}

func (p *Plaid) GetInstitutionInfoById(ctx context.Context, id string) (plaidSdk.InstitutionsGetByIdResponse, error) {
	request := plaidSdk.NewInstitutionsGetByIdRequest(id, []plaidSdk.CountryCode{plaidSdk.COUNTRYCODE_US})
	res, _, err := p.client.PlaidApi.
		InstitutionsGetById(ctx).
		InstitutionsGetByIdRequest(*request).
		Execute()

	return res, err
}

func (p *Plaid) GetAccountsInfo(ctx context.Context, accessToken string) (plaidSdk.AccountsGetResponse, error) {
	request := plaidSdk.NewAccountsGetRequest(accessToken)
	res, _, err := p.client.PlaidApi.
		AccountsGet(ctx).
		AccountsGetRequest(*request).
		Execute()

	return res, err
}

func castPlaidTransactionToDBModel(t plaidSdk.Transaction, account *ent.PlaidInstitutionAccount, builder *ent.TransactionCreate) *ent.TransactionCreate {
	location := t.GetLocation()
	return builder.
		SetFinancialAccountID(t.GetAccountId()).
		SetAmount(t.GetAmount()).
		SetIsoCurrencyCode(t.GetIsoCurrencyCode()).
		SetUnofficialCurrencyCode(t.GetUnofficialCurrencyCode()).
		SetCategory(strings.Join(t.GetCategory(), ",")).
		SetCategoryID(t.GetCategoryId()).
		SetCheckNumber(t.GetCheckNumber()).
		SetDate(t.GetDate()).
		SetDatetime(t.GetDatetime()).
		SetAuthorizedDate(t.GetAuthorizedDate()).
		SetAuthorizedDatetime(t.GetAuthorizedDatetime()).
		SetLocationAddress(location.GetAddress()).
		SetLocationCity(location.GetCity()).
		SetLocationRegion(location.GetRegion()).
		SetLocationPostalCode(location.GetPostalCode()).
		SetLocationLat(location.GetLat()).
		SetLocationLon(location.GetLon()).
		SetLocationStoreNumber(location.GetStoreNumber()).
		SetName(t.GetName()).
		SetMerchantName(t.GetMerchantName()).
		SetPaymentChannel(t.GetPaymentChannel()).
		SetPending(t.GetPending()).
		SetPendingTransactionID(t.GetPendingTransactionId()).
		SetAccountOwner(t.GetAccountOwner()).
		SetTransactionID(t.GetAccountOwner()).
		SetTransactionCode(string(t.GetTransactionCode())).
		SetInstitutionAccount(account)
}

func (p *Plaid) SyncTransactions(ctx context.Context, itemId string) error {
	itemWithAccessToken, err := p.db.PlaidItem.Query().
		Where(plaiditem.ItemID(itemId)).
		Only(ctx)
	if err != nil {
		log.Printf("error fetching plaid item from the db: %v", err)
		return err
	}

	latestTransactionsSync, err := p.db.TransactionSync.Query().
		Where(
			transactionsync.HasItemWith(
				plaiditem.ItemID(itemId),
			),
		).
		Order(
			ent.Desc(transactionsync.FieldCreatedAt),
		).
		First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			latestTransactionsSync = &ent.TransactionSync{}
		} else {
			log.Printf("error fetching transactions sync data: %v", err)
			return err
		}
	}

	accounts, err := itemWithAccessToken.
		QueryInstitution().
		QueryAccounts().
		All(ctx)
	if err != nil {
		log.Printf("error fetching financial accounts: %v", err)
		return err
	}
	accountsMap := make(map[string]*ent.PlaidInstitutionAccount)
	for _, a := range accounts {
		accountsMap[a.AccountID] = a
	}

	cursor := latestTransactionsSync.Cursor
	hasMore := true
	var added []plaidSdk.Transaction
	var updated []plaidSdk.Transaction
	var removed []plaidSdk.RemovedTransaction

	for hasMore {
		request := plaidSdk.NewTransactionsSyncRequest(itemWithAccessToken.AccessToken)

		if cursor != "" {
			request.SetCursor(cursor)
		}

		res, _, err := p.client.PlaidApi.TransactionsSync(ctx).TransactionsSyncRequest(*request).Execute()
		if err != nil {
			log.Printf("failed during fetching a protion of transactions :%v", err)
			return err
		}

		added = append(added, res.GetAdded()...)
		updated = append(updated, res.GetModified()...)
		removed = append(removed, res.GetRemoved()...)

		hasMore = res.GetHasMore()
		cursor = res.GetNextCursor()
	}

	var deleteIts []string
	for _, t := range removed {
		deleteIts = append(deleteIts, t.GetTransactionId())
	}
	for _, t := range updated {
		deleteIts = append(deleteIts, t.GetTransactionId())
	}

	_, err = p.db.Transaction.Delete().
		Where(transaction.TransactionIDIn(deleteIts...)).
		Exec(ctx)
	if err != nil {
		log.Printf("error deleting transactions: %v", err)
		return err
	}

	createBulk := make([]*ent.TransactionCreate, len(added))
	for i, t := range added {
		createBulk[i] = castPlaidTransactionToDBModel(
			t,
			accountsMap[t.GetAccountId()],
			p.db.Transaction.Create(),
		)

		if err := p.UpdateSpendingCategoriesIfNecessary(ctx, t); err != nil {
			log.Printf("failed updating spending categories: %v", err)
			return err
		}
	}

	createUpdatedBulk := make([]*ent.TransactionCreate, len(updated))
	for i, t := range updated {
		createUpdatedBulk[i] = castPlaidTransactionToDBModel(
			t,
			accountsMap[t.GetAccountId()],
			p.db.Transaction.Create(),
		)

		if err := p.UpdateSpendingCategoriesIfNecessary(ctx, t); err != nil {
			log.Printf("failed updating spending categories: %v", err)
			return err
		}
	}

	err = p.db.Transaction.
		CreateBulk(append(createBulk, createUpdatedBulk...)...).
		Exec(ctx)
	if err != nil {
		log.Printf("error saving transaction in the db: %v", err)
		return err
	}

	err = p.db.TransactionSync.Create().
		SetCursor(cursor).
		SetItemID(itemWithAccessToken.ID).
		Exec(ctx)
	if err != nil {
		log.Printf("error saving information about transaction syncto the db: %v", err)
		return err
	}

	if len(added) > 0 || len(updated) > 0 || len(removed) > 0 {
		owner, err := itemWithAccessToken.QueryOwner().Only(ctx)
		if err != nil {
			log.Printf("failed querying plaid item owner: %v", err)
			return err
		}

		newNotification, err := p.notificationService.WithTransactionClient(p.db).
			NewInsightsNotification(ctx, owner.ID)
		if err != nil {
			log.Printf("failed creating insights notification: %v", err)
			return err
		}

		p.notificationService.SendNotificationsToRecipients(
			ctx,
			newNotification,
		)
	}

	return nil
}

func (p *Plaid) UpdateSpendingCategoriesIfNecessary(ctx context.Context, transaction plaidSdk.Transaction) error {
	for _, category := range transaction.GetCategory() {
		if p.spendingCategoryService.CategoryExist(category) {
			return nil
		}

		_, err := p.spendingCategoryService.WithTransactionClient(p.db).Create(
			ctx,
			category,
			transaction.GetCategoryId(),
		)
		if err != nil {
			return err
		}
	}

	return nil
}

// @TODO: move to appropriate place togather will other implementation related to webhooks handling
type WebHookBody struct {
	Environment              string `json:"environment"`
	HistoricalUpdateComplete bool   `json:"historical_update_complete"`
	InitialUpdateComplete    bool   `json:"initial_update_complete"`
	ItemId                   string `json:"item_id"`
	WebhookCode              string `json:"webhook_code"`
	WebhookType              string `json:"webhook_type"`
}

func (p *Plaid) HandleWebHook(ctx context.Context, webHookBody []byte) error {
	var body WebHookBody
	if err := json.Unmarshal(webHookBody, &body); err != nil {
		log.Printf("error parsing webhook body into json: %v", err)
		return err
	}

	log.Printf("body: %v", body)

	return db.WithTransaction(ctx, p.db, func(db *ent.Client) error {
		// @TODO: add checking for webhook type
		err := p.WithTransactionClient(db).SyncTransactions(ctx, body.ItemId)

		if err != nil {
			log.Printf("error synchronising transactions: %v", err)
			return err
		}

		return nil
	})
}

type WebhookVerificationKey struct {
	// The alg member identifies the cryptographic algorithm family used with the key.
	Alg string `json:"alg"`
	// The crv member identifies the cryptographic curve used with the key.
	Crv string `json:"crv"`
	// The kid (Key ID) member can be used to match a specific key. This can be used, for instance, to choose among a set of keys within the JWK during key rollover.
	Kid string `json:"kid"`
	// The kty (key type) parameter identifies the cryptographic algorithm family used with the key, such as RSA or EC.
	Kty string `json:"kty"`
	// The use (public key use) parameter identifies the intended use of the public key.
	Use string `json:"use"`
	// The x member contains the x coordinate for the elliptic curve point.
	X string `json:"x"`
	// The y member contains the y coordinate for the elliptic curve point.
	Y         string `json:"y"`
	CreatedAt int32  `json:"created_at"`
	ExpiredAt int32  `json:"expired_at"`
}

type JWKSKeys struct {
	Keys []WebhookVerificationKey `json:"keys"`
}

func NewWebhookVerificationKeyFromPlaid(input plaidSdk.JWKPublicKey) (WebhookVerificationKey, error) {
	return WebhookVerificationKey{
		Alg:       input.GetAlg(),
		Crv:       input.GetCrv(),
		Kid:       input.GetKid(),
		Kty:       input.GetKty(),
		Use:       input.GetUse(),
		X:         input.GetX(),
		Y:         input.GetY(),
		CreatedAt: input.GetCreatedAt(),
		ExpiredAt: input.GetExpiredAt(),
	}, nil
}

func (p *Plaid) VerifyWebHook(ctx context.Context, webHookBody string, tokenString string) bool {
	// Get Plaid-Verification header
	//tokenString := headers["plaid-signature"]

	// Extract "kid" from the header
	//token, parts, err := new(jwt.Parser).ParseUnverified(
	token, _, err := new(jwt.Parser).ParseUnverified(
		tokenString,
		jwt.MapClaims{},
	)
	if err != nil {
		log.Printf("token parsing without validation failed: %v", err)
		return false
	}

	if token.Method.Alg() != "ES256" {
		log.Print("token signing method does not match")
		return false
	}

	kid := token.Header["kid"].(string)

	keyCache := make(map[string]plaidSdk.JWKPublicKey)

	// Get public key from /webhook_verification_key/get by the use of the extracted key id
	if key, found := keyCache[kid]; !found || key.ExpiredAt.Get() != nil {
		kidsToUpdate := make([]string, 0)

		for k, v := range keyCache {
			if v.ExpiredAt.Get() == nil {
				kidsToUpdate = append(kidsToUpdate, k)
			}
		}

		kidsToUpdate = append(kidsToUpdate, kid)

		for _, k := range kidsToUpdate {
			verificationKeyRequest := plaidSdk.NewWebhookVerificationKeyGetRequest(k)
			verificationKeyRes, _, err := p.client.PlaidApi.
				WebhookVerificationKeyGet(ctx).
				WebhookVerificationKeyGetRequest(*verificationKeyRequest).
				Execute()
			if err != nil {
				log.Printf("error fetching verification key for webhook: %v", err)
				return false
			}

			keyCache[k] = verificationKeyRes.GetKey()
		}
	}

	// Validate the webhook
	if _, found := keyCache[kid]; !found {
		return false
	}

	key := keyCache[kid]

	if key.ExpiredAt.Get() != nil {
		return false
	}

	log.Printf("key ppoints %v, %v ", key.GetX(), key.GetY())

	wrappedKey, _ := NewWebhookVerificationKeyFromPlaid(key)
	jsonKeys, err := json.Marshal(JWKSKeys{Keys: []WebhookVerificationKey{wrappedKey}})
	if err != nil {
		log.Printf("failde during encoding JWKS keys into json: %v", err)
		return false
	}
	jwks, err := keyfunc.NewJSON(jsonKeys)
	if err != nil {
		log.Printf("failed creating JWKS keyfunction: %v", err)
		return false
	}

	//publicKey := new(ecdsa.PublicKey)
	//publicKey.Curve = elliptic.P256()
	//x, err := base64.StdEncoding.DecodeString(key.X)
	//log.Print(x, err)
	//xc := new(big.Int)
	//publicKey.X = xc.SetBytes(x)
	//y, err := base64.StdEncoding.DecodeString(key.Y)
	//log.Print(y, err)
	//yc := new(big.Int)
	//publicKey.Y = yc.SetBytes(y)
	//log.Printf("public key values: %v, %v", publicKey.X, publicKey.Y)

	verifiedToken, err := jwt.Parse(
		tokenString,
		jwks.Keyfunc,
	)
	if err != nil || !verifiedToken.Valid {
		log.Printf("toke is not valid: %v, %v", err, verifiedToken)
		return false
	}

	// Extract the payload of the JWT
	claims := token.Claims.(jwt.MapClaims)
	iat := claims["iat"]

	timeSinceIat := float64(time.Now().Unix()) - iat.(float64)

	// Check if the web hook payload is not older than 5 minutes
	if timeSinceIat > 300 {
		log.Print("time of JWT token issue is more that 5 minutes ago")
		return false
	}

	// Validate the body
	sha256Value := claims["request_body_sha256"]

	sha := sha256.Sum256([]byte(webHookBody))
	sha256Body := hex.EncodeToString(sha[:])

	return sha256Value == sha256Body
}

func (p *Plaid) GetClient() *plaidSdk.APIClient {
	return p.client
}

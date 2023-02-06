package user

import (
	"context"
	"log"
	"wallit/ent"
	"wallit/ent/user"
	"wallit/ent/usernotificationchannelpreferences"
)

type Service struct {
	db *ent.Client
}

func New(db *ent.Client) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) WithTransaction(tx *ent.Tx) *Service {
	return s.WithTransactionClient(tx.Client())
}

func (s *Service) WithTransactionClient(db *ent.Client) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) GetOrCreateProfileWithIdentityProviderId(ctx context.Context, identityProviderId string) (*ent.User, error) {
	existingUser, err := s.db.User.Query().
		Where(
			user.IdentityProviderID(identityProviderId),
		).
		Only(ctx)

	if err != nil && !ent.IsNotFound(err) {
		log.Printf("failed to fetch existing user: %v", err)

		return nil, err
	}

	if existingUser != nil {
		return existingUser, nil
	}

	newUser, err := s.db.User.Create().
		SetIdentityProviderID(identityProviderId).
		Save(ctx)

	if err != nil {
		log.Printf("failed to create a new user profile: %v", err)

		return nil, err
	}

	return newUser, nil
}

func (s *Service) Update(ctx context.Context, userId int, input *ent.UpdateUserInput) (*ent.User, error) {
	return s.db.User.UpdateOneID(userId).
		SetInput(*input).
		Save(ctx)
}

func (s *Service) SaveItemAccessToken(ctx context.Context, userId int, itemId string, accessToken string) (*ent.PlaidItem, error) {
	return s.db.PlaidItem.Create().
		SetOwnerID(userId).
		SetItemID(itemId).
		SetAccessToken(accessToken).
		Save(ctx)
}

func (s *Service) SetSpendingCategories(ctx context.Context, userId int, spendingCategoryIds []int) (*ent.User, error) {
	if err := s.db.User.UpdateOneID(userId).ClearSpendingCategories().Exec(ctx); err != nil {
		log.Printf("failed removing spending categories %v", err)
		return nil, err
	}

	user, err := s.db.User.UpdateOneID(userId).
		AddSpendingCategoryIDs(spendingCategoryIds...).
		Save(ctx)

	if err != nil {
		log.Printf("failed updating user pending categories: %v", err)
		return nil, err
	}

	return user, nil
}

func (s *Service) SetNotificationChannels(ctx context.Context, userId int, notificationChannels []usernotificationchannelpreferences.Chanel) (*ent.User, error) {
	if err := s.db.User.UpdateOneID(userId).ClearNotificationChannels().Exec(ctx); err != nil {
		log.Printf("failed clearing user notification channels: %v", err)
		return nil, err
	}

	notificationChannelCreateBulk := make([]*ent.UserNotificationChannelPreferencesCreate, len(notificationChannels))
	for i, n := range notificationChannels {
		notificationChannelCreateBulk[i] = s.db.UserNotificationChannelPreferences.
			Create().
			SetChanel(n).
			SetChanelUsersID(userId)
	}

	err := s.db.UserNotificationChannelPreferences.
		CreateBulk(notificationChannelCreateBulk...).
		Exec(ctx)
	if err != nil {
		log.Printf("failed saving user notification preferences: %v", err)
	}

	return s.db.User.Query().Where(user.ID(userId)).Only(ctx)
}

func (s *Service) SetOfferFrequency(ctx context.Context, userId int, frequency user.OfferFrequency) (*ent.User, error) {
	return s.db.User.UpdateOneID(userId).
		SetOfferFrequency(frequency).
		Save(ctx)
}

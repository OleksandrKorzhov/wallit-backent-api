package auth

import (
	"context"
	"fmt"
	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
	"wallit/internal/config"
)

type CustomClaims struct {
	Scope string `json:"scope"`
}

func (c *CustomClaims) Validate(ctx context.Context) error {
	return nil
}

type jwtClaims struct{}

func NewTokenValidator() func(next http.Handler) http.Handler {
	issuer := config.Get[string]("AUTH0_DOMAIN")
	audience := config.Get[string]("AUTH0_AUDIENCE")

	issuerUrl, err := url.Parse("https://" + issuer + "/")
	if err != nil {
		log.Fatalf("failed creating issuer url: %v", err)
	}

	jwksProvider := jwks.NewCachingProvider(issuerUrl, time.Minute*5)

	jwtValidator, err := validator.New(
		jwksProvider.KeyFunc,
		validator.RS256,
		issuerUrl.String(),
		[]string{audience},
		validator.WithCustomClaims(
			func() validator.CustomClaims {
				return &CustomClaims{}
			}),
		validator.WithAllowedClockSkew(time.Minute),
	)
	if err != nil {
		log.Fatalf("failed creating JWT validator: %v", err)
	}

	//jwtmiddleware.New()

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString, err := jwtmiddleware.AuthHeaderTokenExtractor(r)
			if err != nil {
				fmt.Print(fmt.Errorf("error extracting JWT from auth header: %v", err))
				next.ServeHTTP(w, r)
				return
			}

			validationResult, err := jwtValidator.ValidateToken(r.Context(), tokenString)
			if err != nil {
				fmt.Print(fmt.Errorf("failed to validate access token: %v", err))
				next.ServeHTTP(w, r)
				return
			}

			log.Printf("%v", validationResult)

			validatedTokenClaims, ok := validationResult.(*validator.ValidatedClaims)

			log.Printf("%v", validatedTokenClaims.CustomClaims)
			if !ok {
				fmt.Print(fmt.Errorf("failed cast JWT validation result to token claims: %v", err))
				next.ServeHTTP(w, r)
				return
			}

			ctxWithClaims := context.WithValue(r.Context(), &jwtClaims{}, validatedTokenClaims.CustomClaims)

			next.ServeHTTP(w, r.Clone(ctxWithClaims))
		})
	}
}

func FromContext(ctx context.Context) *CustomClaims {
	value, ok := ctx.Value(&jwtClaims{}).(*CustomClaims)

	if !ok {
		return &CustomClaims{}
	}

	return value
}

func (c *CustomClaims) HasScope(scope string) bool {
	scopes := strings.Fields(c.Scope)

	for _, s := range scopes {
		if s == scope {
			return true
		}
	}

	return false
}

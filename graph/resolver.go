package graph

import (
	"github.com/99designs/gqlgen/graphql"
	"wallit/ent"
	"wallit/internal/discount_offer"
	"wallit/internal/merchant"
	"wallit/internal/notification"
	"wallit/internal/plaid"
	"wallit/internal/spending_category"
	"wallit/internal/user"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	db                      *ent.Client
	plaid                   *plaid.Plaid
	userService             *user.Service
	spendingCategoryService *spending_category.Service
	notificationService     *notification.Service
	discountOfferService    *discount_offer.Service
	merchantService         *merchant.Service
}

func NewSchema(
	db *ent.Client,
	plaidInst *plaid.Plaid,
	userService *user.Service,
	spendingCategoryService *spending_category.Service,
	notificationService *notification.Service,
	discountOfferService *discount_offer.Service,
	merchantService *merchant.Service,
) graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers: &Resolver{
			db:                      db,
			plaid:                   plaidInst,
			userService:             userService,
			spendingCategoryService: spendingCategoryService,
			notificationService:     notificationService,
			discountOfferService:    discountOfferService,
			merchantService:         merchantService,
		},
	})
}

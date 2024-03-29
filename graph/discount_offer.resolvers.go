package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"wallit/ent"
)

// CreateDiscountOffer is the resolver for the createDiscountOffer field.
func (r *mutationResolver) CreateDiscountOffer(ctx context.Context, input ent.CreateDiscountOfferInput) (*ent.DiscountOffer, error) {
	client := ent.FromContext(ctx)
	//client, err := r.db.Tx(ctx)
	//if err != nil {
	//	log.Printf("failed to start transaction during discount offer creation: %v", err)
	//	return nil, err
	//}

	return r.discountOfferService.
		//WithTransactionClient(client.Client()).
		WithTransactionClient(client).
		NewDiscountOffer(ctx, input)
}

// DiscountOffers is the resolver for the discountOffers field.
func (r *queryResolver) DiscountOffers(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.DiscountOfferOrder, where *ent.DiscountOfferWhereInput) (*ent.DiscountOfferConnection, error) {
	return r.db.DiscountOffer.Query().
		Paginate(
			ctx,
			after,
			first,
			before,
			last,
			ent.WithDiscountOfferFilter(where.Filter),
			ent.WithDiscountOfferOrder(orderBy),
		)
}

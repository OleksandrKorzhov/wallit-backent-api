package merchant

import (
	"context"
	"wallit/ent"
)

type Service struct {
	db *ent.Client
}

func New(db *ent.Client) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) NewMerchant(ctx context.Context, input ent.CreateMerchantInput) (*ent.Merchant, error) {
	return s.db.Merchant.Create().
		SetInput(input).
		Save(ctx)
}

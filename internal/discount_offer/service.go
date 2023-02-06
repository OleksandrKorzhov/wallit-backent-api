package discount_offer

import (
	"context"
	"log"
	"wallit/ent"
	"wallit/internal/notification"
)

type Service struct {
	db                  *ent.Client
	notificationService *notification.Service
}

func New(db *ent.Client, notificationService *notification.Service) *Service {
	return &Service{
		db:                  db,
		notificationService: notificationService,
	}
}

func (s *Service) WithTransactionClient(db *ent.Client) *Service {
	return &Service{
		db:                  db,
		notificationService: s.notificationService,
	}
}

func (s *Service) NewDiscountOffer(ctx context.Context, input ent.CreateDiscountOfferInput) (*ent.DiscountOffer, error) {
	offer, err := s.db.DiscountOffer.Create().
		SetInput(input).
		Save(ctx)
	if err != nil {
		log.Printf("failed creating discount: %v", err)
		return nil, err
	}

	txNotificationsService := s.notificationService.WithTransactionClient(s.db)

	notifications, err := txNotificationsService.NewOfferNotification(ctx, offer)
	if err != nil {
		log.Printf("failed creating discount notifications: %v", err)
		return nil, err
	}

	txNotificationsService.SendNotificationsToRecipients(ctx, notifications...)

	return offer, nil
}

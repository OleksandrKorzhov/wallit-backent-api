package notification

import (
	"context"
	"log"
	"sync"
	"time"
	"wallit/ent"
	"wallit/ent/notification"
)

type Service struct {
	db         *ent.Client
	recipients map[int]chan *ent.Notification
	mu         sync.Mutex
}

func New(db *ent.Client) *Service {
	return &Service{
		db:         db,
		recipients: make(map[int]chan *ent.Notification),
	}
}

func (s *Service) WithTransaction(tx *ent.Tx) *Service {
	return s.WithTransactionClient(tx.Client())
}

func (s *Service) WithTransactionClient(db *ent.Client) *Service {
	//return s
	return &Service{
		db:         db,
		recipients: s.recipients,
		mu:         s.mu,
	}
}

func (s *Service) NewOfferNotification(ctx context.Context, discountOffer *ent.DiscountOffer) ([]*ent.Notification, error) {
	eligibleUsers, err := discountOffer.DiscountEligibleUsers(ctx)
	if err != nil {
		log.Printf("failed fetching eligible users for discount offer during notification sending: %v", err)
		return nil, err
	}

	bulkCreate := make([]*ent.NotificationCreate, len(eligibleUsers))

	for i, user := range eligibleUsers {
		bulkCreate[i] = s.db.Notification.Create().
			SetType(notification.TypeOffer).
			SetNotificationRecipientID(user.ID).
			SetNotificationDiscountOfferID(discountOffer.ID)
	}

	return s.db.Notification.CreateBulk(bulkCreate...).Save(ctx)
}

func (s *Service) NewInsightsNotification(ctx context.Context, ownerId int) (*ent.Notification, error) {
	return s.db.Notification.Create().
		SetType(notification.TypeInsights).
		SetNotificationRecipientID(ownerId).
		Save(ctx)
}

func (s *Service) MarkAsCheckedInApp(ctx context.Context, ids ...int) error {
	return s.db.Notification.Update().
		Where(notification.IDIn(ids...)).
		SetCheckedInApp(true).
		Exec(ctx)
}

func (s *Service) NewNotificationRecipient(recipientId int) <-chan *ent.Notification {
	channel := make(chan *ent.Notification)

	s.mu.Lock()
	s.recipients[recipientId] = channel
	s.mu.Unlock()

	return channel
}

func (s *Service) RemoveNotificationRecipient(recipientId int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	recipient, ok := s.recipients[recipientId]
	if !ok {
		log.Printf("Recipient '%v' is not found", recipient)
		return
	}

	close(s.recipients[recipientId])

	delete(s.recipients, recipientId)
}

func (s *Service) SendNotificationsToRecipients(ctx context.Context, notifications ...*ent.Notification) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, n := range notifications {
		// @TODO: change to notification recipient
		recipient, err := n.NotificationRecipient(ctx)
		if err != nil {
			log.Printf("failed getting notification owner during notification send: %v", err)
			continue
		}

		recipientChan, ok := s.recipients[recipient.ID]
		if !ok {
			log.Printf("Recipient '%v' is not found! Cannot send the notification!", recipient.ID)
			return
		}

		select {
		case recipientChan <- n:
			log.Printf("sent notification %v of type %v to recipient %v", n.ID, n.Type, recipient.ID)
		case <-time.After(time.Second * 10):
			log.Printf("notification sending to recipient %v timedout", recipient.ID)
		}
	}
}

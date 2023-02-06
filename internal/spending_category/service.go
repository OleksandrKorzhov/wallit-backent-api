package spending_category

import (
	"context"
	"log"
	"wallit/ent"
)

type Service struct {
	db         *ent.Client
	categories map[string]int
}

func New(db *ent.Client) *Service {
	s := &Service{
		db:         db,
		categories: make(map[string]int),
	}

	s.LoadIntoCache()

	return s
}

func (s *Service) WithTransactionClient(db *ent.Client) *Service {
	return &Service{
		db:         db,
		categories: s.categories,
	}
}

func (s *Service) addCategoryIntoCache(category *ent.SpendingCategory) {
	s.categories[category.Name] = category.ID
}

func (s *Service) LoadIntoCache() {
	categories, err := s.db.SpendingCategory.Query().All(context.Background())
	if err != nil {
		log.Printf("failed loading categories: %v", err)
		return
	}

	for _, category := range categories {
		s.addCategoryIntoCache(category)
	}
}

func (s *Service) Create(ctx context.Context, name string, categoryId string) (*ent.SpendingCategory, error) {
	category, err := s.db.SpendingCategory.Create().
		SetName(name).
		SetCategoryID(categoryId).
		Save(ctx)

	if err != nil {
		log.Printf("failed creating new category: %v", err)
		return nil, err
	}

	s.addCategoryIntoCache(category)

	return category, nil
}

func (s *Service) CategoryExist(categoryName string) bool {
	_, exist := s.categories[categoryName]

	return exist
}

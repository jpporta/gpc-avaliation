package services

import (
	"context"
	"database/sql"

	"github.com/jpporta/gpc-avaliation/internal/repository"
)

func (s *Service) CreateCategory(ctx context.Context, name, description string) (int64, error) {
	id, err := s.queries.CreateCategory(ctx, repository.CreateCategoryParams{
		Name: name,
		Description: sql.NullString{
			String: description,
			Valid:  description != "",
		},
	})
	if err != nil {
		return 0, err
	}
	return id, nil
}

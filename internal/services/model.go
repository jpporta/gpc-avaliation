package services

import (
	"context"
	"database/sql"

	"github.com/jpporta/gpc-avaliation/internal/repository"
)

func (s *Service) CreateModel(ctx context.Context, name, author, description string, categoryId int64) (int64, error) {
	id, err := s.queries.CreateModel(ctx, repository.CreateModelParams{
		Name:   name,
		Author: author,
		Description: sql.NullString{
			String: description,
			Valid:  description != "",
		},
		CategoryID: categoryId,
	})
	if err != nil {
		return 0, err
	}
	return id, nil
}

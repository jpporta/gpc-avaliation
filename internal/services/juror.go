package services

import (
	"context"

	"github.com/jpporta/gpc-avaliation/internal/repository"
	"github.com/jpporta/gpc-avaliation/internal/utils"
)

func (s *Service) CreateJuror(ctx context.Context, name string) (int64, error) {
	slug := utils.RandomSlug()
	id, err := s.queries.CreateJuror(ctx, repository.CreateJurorParams{
		Slug: slug,
		Name: name,
	})
	if err != nil {
		return 0, err
	}
	return id, nil
}

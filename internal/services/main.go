package services

import (
	"database/sql"

	_ "github.com/glebarez/go-sqlite"
	"github.com/jpporta/gpc-avaliation/internal/repository"
	"github.com/jpporta/gpc-avaliation/internal/utils"
)

type Service struct {
	queries *repository.Queries
}

func New() *Service {
	db := utils.Must(sql.Open("sqlite",
		"./gpc.db"),
	)

	q := repository.New(db)
	return &Service{
		queries: q,
	}
}

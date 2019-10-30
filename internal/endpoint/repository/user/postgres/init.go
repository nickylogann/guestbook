package postgres

import (
	"github.com/tokopedia/sqlt"

	"github.com/nickylogan/guestbook/internal/endpoint/repository/user"
	"github.com/nickylogan/guestbook/internal/pkg/config"
)

type postgresRepository struct {
	db        *sqlt.DB
	config    *config.RuntimeConfig
	tableName string
}

func NewRepository(db *sqlt.DB, config *config.RuntimeConfig, tableName string) user.Repository {
	return &postgresRepository{
		db:        db,
		config:    config,
		tableName: tableName,
	}
}
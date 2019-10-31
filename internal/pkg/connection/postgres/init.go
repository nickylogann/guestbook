package postgres

import (
	"fmt"

	// Import postgres driver
	_ "github.com/lib/pq"
	"github.com/nickylogan/guestbook/internal/pkg/utils/config"
	"github.com/tokopedia/sqlt"
)

// NewConnection creates a new postgres connection
func NewConnection(cfg config.PostgresConfig) (client *sqlt.DB, err error) {
	DSN := fmt.Sprintf("host=%s port=5432 dbname=%s user=%s password=%s sslmode=disable", cfg.Host, cfg.Name, cfg.User, cfg.Password)

	client, err = sqlt.Open("postgres", DSN)
	if err != nil {
		return
	}
	client.SetMaxOpenConnections(100)
	return
}

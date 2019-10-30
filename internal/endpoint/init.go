package endpoint

import (
	"log"

	"github.com/nickylogan/guestbook/internal/pkg/config"
)

// Run runs the app
func Run() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Println("[error][endpoint][Run] failed to init config:", err)
		return
	}
	// TODO: add init logic
	_ = cfg
}
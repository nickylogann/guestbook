package app

import (
	"fmt"
	"log"
	"os"

	"github.com/labstack/echo"

	"github.com/nickylogan/guestbook/internal/app/delivery/http"
	redisRepo "github.com/nickylogan/guestbook/internal/app/repository/visitor/redis"
	. "github.com/nickylogan/guestbook/internal/app/usecase/user/functions"
	. "github.com/nickylogan/guestbook/internal/app/usecase/visitor/functions"
	"github.com/nickylogan/guestbook/internal/pkg/connection/redis"

	"github.com/joho/godotenv"

	postgresUserRepo "github.com/nickylogan/guestbook/internal/app/repository/user/postgres"
	"github.com/nickylogan/guestbook/internal/pkg/connection/postgres"
	"github.com/nickylogan/guestbook/internal/pkg/utils/config"
)

// Run runs the app
func Run() {
	err := godotenv.Load("./configs/.env.development")
	if err != nil {
		log.Println(err)
		return
	}
	cfg := config.NewConfig()

	dbConn, err := postgres.NewConnection(cfg.Database.Postgres)
	if err != nil {
		log.Fatal(err, cfg.Database.Postgres)
	}
	redisConn := redis.NewConnection(cfg.Database.Redis)

	e := echo.New()
	userRepo := postgresUserRepo.NewPostgresRepository(dbConn, &cfg, os.Getenv("PG_TABLE"))
	userUseCase := NewUserUseCase(userRepo)

	visitorRepo := redisRepo.NewRedisRepository(redisConn, os.Getenv("REDIS_KEY"))
	visitorUseCase := NewVisitorUseCase(visitorRepo)

	http.NewUserHandler(e, userUseCase, visitorUseCase)

	fmt.Println("Server started at :8080")
	err = e.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

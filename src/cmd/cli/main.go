package main

import (
	"bytesbanana/realworld-go-echo/src/internal/adapter/db"
	"bytesbanana/realworld-go-echo/src/internal/adapter/handler"
	v "bytesbanana/realworld-go-echo/src/internal/adapter/validator"
	"bytesbanana/realworld-go-echo/src/internal/core/service"
	"encoding/json"
	"log"

	"github.com/go-playground/validator/v10"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"

	_ "github.com/lib/pq"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()

	e := echo.New()
	e.Validator = &v.CustomValidator{Validator: validator.New()}

	e.Use(middleware.RequestLoggerWithConfig(
		middleware.RequestLoggerConfig{
			LogURI:    true,
			LogStatus: true,
			LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
				logger.Info("request",
					zap.String("URI", v.URI),
					zap.Int("status", v.Status),
				)

				return nil
			},
		},
	))

	api := e.Group("/api")
	dsn := "dbname=realworld host=localhost port=5432 user=postgres password=password sslmode=disable"

	logger.Info("connecting to database")
	dbx, err := sqlx.Open("postgres", dsn)
	if err != nil {
		logger.Fatal("failed to connect to database", zap.Error(err))
	}
	defer dbx.Close()

	ur := db.NewUserRepository(dbx)
	us := service.NewUserService(ur)
	h := handler.New(us)

	api.POST("/users", h.CreateUser)

	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		logger.Fatal("failed to marshal routes", zap.Error(err))
	}
	logger.Info(string(data))

	e.Logger.Fatal(e.Start(":1323"))

}

package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"io"
	"log"
	"net/http"
	"wallit/ent"
	"wallit/internal/config"
	"wallit/internal/notification"
	"wallit/internal/plaid"
	"wallit/internal/spending_category"
)

func main() {
	config.Init()
	// @TODO: move DB setup to an independent functionality
	dbHost := config.Get[string]("DB_HOST")
	dbPort := config.Get[string]("DB_PORT")
	dbUser := config.Get[string]("DB_USER")
	dbPassword := config.Get[string]("DB_PASSWORD")
	dbName := config.Get[string]("DB_NAME")

	var entOptions []ent.Option
	entOptions = append(entOptions, ent.Debug()) // @TODO: add condition for enabling/disabling database debug

	client, err := ent.Open("postgres", fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPassword))
	if err != nil {
		log.Fatalf("Failed to open connection to postgres: %v", err)
	}

	defer client.Close()

	//if err = client.Schema.Create(context.Background(), migrate.WithGlobalUniqueID(true)); err != nil {
	//	log.Fatalf("Failed creating schema resources: %v", err)
	//}

	spendingCategoriesService := spending_category.New(client)
	notificationsService := notification.New(client)
	plaidService := plaid.New(
		spendingCategoriesService,
		notificationsService,
		client,
	)
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.POST("/webhook/plaid", func(e echo.Context) error {
		request := e.Request()
		body, err := io.ReadAll(request.Body)
		if err != nil {
			return err
		}

		webhookValid := plaidService.VerifyWebHook(
			request.Context(),
			string(body),
			request.Header.Get("plaid-verification"),
		)

		if !webhookValid {
			return errors.New("webhook validation failed")
		}

		if err := plaidService.HandleWebHook(context.Background(), body); err != nil {
			return e.String(http.StatusInternalServerError, "Error")
		}

		return e.String(http.StatusOK, "Ok")
	})

	log.Printf("Starting server on :4000 port")

	if err := e.Start(":4000"); err != nil {
		log.Fatal(err)
	}
}

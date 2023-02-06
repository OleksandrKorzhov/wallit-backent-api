package main

import (
	"context"
	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
	"net/http"
	"time"
	"wallit/ent"
	"wallit/ent/migrate"
	"wallit/internal/discount_offer"
	"wallit/internal/merchant"
	"wallit/internal/notification"
	"wallit/internal/pkg/auth"
	"wallit/internal/spending_category"

	//"entgo.io/ent/entc/integration/ent"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	//"github.com/labstack/echo/v4"
	//"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"log"
	"wallit/graph"
	"wallit/internal/config"
	"wallit/internal/plaid"
	"wallit/internal/user"
)

func main() {
	config.Init()
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

	if err = client.Schema.Create(context.Background(), migrate.WithGlobalUniqueID(true)); err != nil {
		log.Fatalf("Failed creating schema resources: %v", err)
	}

	notificationService := notification.New(client)
	userService := user.New(client)
	spendingCategoryService := spending_category.New(client)
	merchantService := merchant.New(client)
	discountOfferService := discount_offer.New(client, notificationService)
	plaidInst := plaid.New(
		spendingCategoryService,
		notificationService,
		client,
	)

	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)
	router.Use(cors.AllowAll().Handler)
	//e := echo.New()

	//e.Use(middleware.Recover())
	//e.Use(middleware.Logger())
	//e.Use(middleware.CORS())
	router.Use(auth.NewTokenValidator())
	//router.Use(func(next http.Handler) http.Handler {
	//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//		authScrop := auth.FromContext(r.Context())
	//
	//		fmt.Printf("auth scropes: %v", authScrop)
	//
	//		next.ServeHTTP(w, r)
	//	})
	//})

	//srv := handler.NewDefaultServer(
	srv := handler.New(
		graph.NewSchema(
			client,
			plaidInst,
			userService,
			spendingCategoryService,
			notificationService,
			discountOfferService,
			merchantService,
		),
	)
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})
	srv.SetQueryCache(lru.New(1000))
	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})
	// add ws transport configured by ourselves
	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			//ReadBufferSize:  1024,
			//WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				// add checking origin logic to decide return true or false
				return true
			},
		},
		KeepAlivePingInterval: 10 * time.Second,
	})
	//srv.AddTransport(&transport.Websocket{})
	srv.Use(entgql.Transactioner{
		TxOpener: client,
	})

	router.Handle("/query", srv)
	router.Handle("/playground", playground.Handler("GraphQL", "/query"))

	fmt.Printf("Connect to http://localhost:3200/playground for GraphQL playground")
	log.Fatal(http.ListenAndServe(":3200", router))
}

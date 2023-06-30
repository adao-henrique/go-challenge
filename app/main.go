package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/adao-henrique/go-challenge/app/config"
	http_api "github.com/adao-henrique/go-challenge/app/http"
	handler_account "github.com/adao-henrique/go-challenge/app/http/account"
	handler_login "github.com/adao-henrique/go-challenge/app/http/login"
	handler_transfer "github.com/adao-henrique/go-challenge/app/http/tranfer"
	_ "github.com/adao-henrique/go-challenge/docs"
	usecase_account "github.com/adao-henrique/go-challenge/domain/account"
	service_loin "github.com/adao-henrique/go-challenge/domain/login"
	usecase_transfer "github.com/adao-henrique/go-challenge/domain/transfer"
	repository_accouunt "github.com/adao-henrique/go-challenge/gateways/postgres/account"
	repository_transfer "github.com/adao-henrique/go-challenge/gateways/postgres/transfer"
	httpSwagger "github.com/swaggo/http-swagger/v2"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
)

func main() {

	log.Println("Init start application")
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading .env file ", err)
	}

	// log.Println("Init migration")
	// Migrate(config.Infra.Postgres.Address())

	log.Println("Init DB connection")
	conn, err := ConnectDB(config.Infra.Postgres.Address(), config.Infra.Postgres.PoolMinSize, config.Infra.Postgres.PoolMaxSize)
	if err != nil {
		log.Fatal("setuping postgres: %w", err)
	}

	log.Println("Init create repository")
	accountRepository := repository_accouunt.NewRepository(conn)
	transferRepository := repository_transfer.NewRepository(conn)

	log.Println("Init create usecases")
	accounttUseCase := usecase_account.NewUseCase(accountRepository)
	transferUseCase := usecase_transfer.NewUseCase(transferRepository, accounttUseCase)
	loginService := service_loin.NewService(accounttUseCase)

	log.Println("Init create handlers")
	acccountHandler := handler_account.NewHandler(accounttUseCase)
	transferHandler := handler_transfer.NewHandler(transferUseCase)
	loginHandler := handler_login.NewHandler(loginService)

	r := chi.NewRouter()

	log.Println("Init create API")
	http_api.NewApi(r, acccountHandler, transferHandler, loginHandler)
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), //The url pointing to API definition
	))

	log.Println("Start apllication")
	http.ListenAndServe(":8080", r)

}

func Migrate(addr string) {
	log.Println("addr: ", addr)
	// db, err := sql.Open("postgres", addr)
	// if err != nil {
	// 	log.Fatalf("Error coneect DB to migrate ", err)
	// }

	// defer db.Close()

	// // connect to check connection
	// if err := db.Ping(); err != nil {
	// 	log.Fatal("Failed to ping db ", err)
	// }

	m, err := migrate.New(
		"file://../gateways/postgres/migrations",
		addr,
	)
	if err != nil {
		log.Fatalf("Error create migrate instance ", err)
	}

	m.Up()
}

func ConnectDB(addr string, minConn, maxConn int32) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(addr)
	if err != nil {
		return nil, fmt.Errorf("parsing pgxpool config: %w", err)
	}

	// The defaults are located on top of pgxpool.pool.go
	config.MaxConns = maxConn
	config.MinConns = minConn

	pgxConn, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("creating new pgxpool: %w", err)
	}

	return pgxConn, nil
}

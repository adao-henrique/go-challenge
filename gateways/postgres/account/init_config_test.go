package account

// import (
// 	"context"
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strings"
// 	"testing"
// 	"time"

// 	"github.com/golang-migrate/migrate"
// 	"github.com/golang-migrate/migrate/database/postgres"
// 	"github.com/jackc/pgx/v5"
// 	"github.com/joho/godotenv"
// 	"github.com/ory/dockertest/v3"
// 	"github.com/ory/dockertest/v3/docker"
// )

// var (
// 	db          *pgx.Conn
// 	password    string
// 	user        string
// 	dbName      string
// 	hostAndPort string
// )

// func TestMain(m *testing.M) {
// 	log.Println("Do stuff BEFORE the tests!")
// 	cwd, err := os.Getwd()
// 	log.Println(cwd, err)

// 	err = godotenv.Load("../../../.env")

// 	if err != nil {
// 		log.Fatalf("Error loading .env file")
// 	}

// 	password = os.Getenv("POSTGRES_PASSWORD")
// 	user = os.Getenv("POSTGRES_USER")
// 	dbName = os.Getenv("POSTGRES_DBNAME")
// 	ctx := context.Background()

// 	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
// 	pool, err := dockertest.NewPool("")
// 	if err != nil {
// 		log.Fatalf("Could not construct pool: %s", err)
// 	}

// 	err = pool.Client.Ping()
// 	if err != nil {
// 		log.Fatalf("Could not connect to Docker: %s", err)
// 	}

// 	// pulls an image, creates a container based on it and runs it
// 	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
// 		Repository: "postgres",
// 		Tag:        "11",
// 		Env: []string{
// 			"POSTGRES_PASSWORD=" + password,
// 			"POSTGRES_USER=" + user,
// 			"POSTGRES_DB=" + dbName,
// 			"listen_addresses = '*'",
// 		},
// 	}, func(config *docker.HostConfig) {
// 		// set AutoRemove to true so that stopped container goes away by itself
// 		config.AutoRemove = true
// 		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
// 	})
// 	if err != nil {
// 		log.Fatalf("Could not start resource: %s", err)
// 	}

// 	hostAndPort = resource.GetHostPort("5432/tcp")
// 	databaseUrl := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", user, password, hostAndPort, dbName)

// 	log.Println("Connecting to database on url: ", databaseUrl)

// 	resource.Expire(120) // Tell docker to hard kill the container in 120 seconds

// 	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
// 	pool.MaxWait = 120 * time.Second
// 	if err = pool.Retry(func() error {

// 		db, err = pgx.Connect(ctx, databaseUrl)
// 		if err != nil {
// 			return err
// 		}
// 		return db.Ping(ctx)
// 	}); err != nil {
// 		log.Fatalf("Could not connect to docker: %s", err)
// 	}
// 	MigrateDb(databaseUrl)

// 	//Run tests
// 	code := m.Run()

// 	log.Println("Do stuff AFTER the tests!")

// 	// You can't defer this because os.Exit doesn't care for defer
// 	if err := pool.Purge(resource); err != nil {
// 		log.Fatalf("Could not purge resource: %s", err)
// 	}

// 	os.Exit(code)
// }

// func MigrateDb(dbURI string) error {
// 	hostPort := strings.Split(hostAndPort, ":")
// 	dbURI = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", hostPort[0], hostPort[1], user, password, dbName)
// 	conn, err := sql.Open("postgres", dbURI)
// 	if err != nil {
// 		panic(err)
// 	}
// 	err = conn.Ping()
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("database connection successful")

// 	// driver, err := postgres.WithInstance(conn, &postgres.Config{})
// 	// if err != nil {
// 	// 	panic(err)
// 	// }

// 	driver, _ := postgres.WithInstance(conn, &postgres.Config{})
// 	m, _ := migrate.NewWithDatabaseInstance(
// 		"file://gateways/postgres/migrations",
// 		fmt.Sprintf("postgres://%s:%s@/%s", user, password, dbName),
// 		driver,
// 	)

// 	// m, err := migrate.NewWithDatabaseInstance(
// 	// 	"file://gateways/postgres/migrations",
// 	// 	"postgres",
// 	// 	driver)
// 	if err != nil {
// 		return err
// 	}
// 	err = m.Up()
// 	if err != nil && err != migrate.ErrNoChange {
// 		panic(err)
// 	}

// 	fmt.Println("database migration successful")
// 	return nil
// }

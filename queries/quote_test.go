package queries

import (
	"database/sql"
	"log"
	"os"
	"testing"
	protos "github.com/jeebster/go-grpc-server-quotes/protos"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func TestMain(m *testing.M) {
	initGoDotEnv()
	code := m.Run() // run the test suite
	os.Exit(code)
}

func TestGetQuotes(t *testing.T) {
	// setup
	dbConnection := initDbConnection()
	defer dbConnection.Close()
	createTable(dbConnection)
	seedTable(dbConnection)

	// tests
	limit := 3
	quotes := GetQuotes(dbConnection, limit)
	assert.Equal(t, len(quotes), limit, "results count should be relative to the limit argument")

	// teardown
	truncateTable(dbConnection)
}

func createTable(dbConnection *sql.DB) {
	dbConnection.Exec(`CREATE TABLE IF NOT EXISTS quotes (
		id         SERIAL                   NOT NULL PRIMARY KEY,
		body       TEXT                     NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT  CURRENT_TIMESTAMP,
		updated_at TIMESTAMP WITH TIME ZONE DEFAULT  CURRENT_TIMESTAMP
	)`)
}

fn seedTable(dbConnection *sql.DB) {
	dbConnection.Exec(`INSERT INTO quotes (body) VALUES
		('this is the first quote'),
		('this is the second quote'),
		('this is the third quote')
	`)
}

func initDbConnection(dbURL string) *sql.DB {
	dbURL := os.Getenv("TEST_DATABASE_URL")
	db, err := sql.Open("postgres", dbURL)

	if err != nil {
		log.Fatal("Error connecting to the test database")
	}

	return db
}

func initGoDotEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func truncateTable(dbConnection *sql.DB) {
	dbConnection.Exec("TRUNCATE TABLE quotes")
	dbConnection.Exec("RESTART IDENTITY")
}

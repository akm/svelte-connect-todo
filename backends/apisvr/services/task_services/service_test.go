package taskservices

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-testfixtures/testfixtures/v3"
)

var (
	db       *sql.DB
	fixtures *testfixtures.Loader
)

func TestMain(m *testing.M) {
	var err error

	// Open connection to the test database.
	// Do NOT import fixtures in a production database!
	// Existing data would be deleted.
	dsn := os.Getenv("TEST_DB_DSN")
	log.Printf("TEST_DB_DSN: %s", dsn)
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("unable to open database: %v", err)
	}

	fixtureDir := os.Getenv("TEST_FIXTURE_DIR")
	log.Printf("TEST_FIXTURE_DIR: %s", fixtureDir)
	fixtures, err = testfixtures.New(
		testfixtures.Database(db),          // You database connection
		testfixtures.Dialect("mysql"),      // Available: "postgresql", "timescaledb", "mysql", "mariadb", "sqlite" and "sqlserver"
		testfixtures.Directory(fixtureDir), // The directory containing the YAML files
		testfixtures.SkipResetSequences(),  // Disable the execution of the SQL command that resets the sequences
	)
	if err != nil {
		log.Fatalf("unable to load fixtures: %v", err)
	}

	os.Exit(m.Run())
}

func prepareTestDatabase(t *testing.T) {
	if err := fixtures.Load(); err != nil {
		t.Fatalf("unable to load fixtures: %v", err)
	}
}

func TestTaskServiceList(t *testing.T) {
	prepareTestDatabase(t)

	t.Error("not implemented")
}

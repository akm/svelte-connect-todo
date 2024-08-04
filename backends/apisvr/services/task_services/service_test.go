package taskservices

import (
	"database/sql"
	"log"
	"os"
	"testing"

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
	db, err = sql.Open("mysql", os.Getenv("TEST_DB_DSN"))
	if err != nil {

	}

	fixtures, err = testfixtures.New(
		testfixtures.Database(db),                             // You database connection
		testfixtures.Dialect("mysql"),                         // Available: "postgresql", "timescaledb", "mysql", "mariadb", "sqlite" and "sqlserver"
		testfixtures.Directory(os.Getenv("TEST_FIXTURE_DIR")), // The directory containing the YAML files
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
}

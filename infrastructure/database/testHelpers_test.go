package database

import (
	"bytes"
	"database/sql"
	"github.com/joho/godotenv"
	"github.com/rhodeon/moviescreen/internal/testhelpers"
	"os"
	"os/exec"
	"testing"
)

// newTestDb generates a database for testing.
// It creates and returns a function for dropping concerned tables whenever called.
func newTestDb(t *testing.T) (*sql.DB, func()) {
	t.Helper()

	// load environment variables from dotenv file
	err := godotenv.Load("./../../.env")
	if err != nil {
		t.Fatal(err)
	}

	dsn := os.Getenv("TEST_DB_DSN")
	db, err := sql.Open("postgres", dsn)
	testhelpers.AssertFatalError(t, err)

	// create tables
	execShellScript(t, "./testdata/setup.sh")

	// populate tables
	execSqlScript(t, db, "./testdata/populate.sql")

	// return function to rollback database changes
	return db, func() {
		execShellScript(t, "./testdata/teardown.sh")
		defer db.Close()
	}
}

// execSqlScript is a helper function to execute shell commands
// in the file at the given scriptPath.
func execShellScript(t *testing.T, scriptPath string) {
	t.Helper()

	cmd := exec.Command("/bin/bash", scriptPath)

	// override standard error buffer of the command
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		t.Fatal(err.Error() + ":\n" + stderr.String())
	}
}

// execSqlScript is a helper function to execute SQL commands
// in the file at the given scriptPath.
func execSqlScript(t *testing.T, db *sql.DB, scriptPath string) {
	t.Helper()

	script, err := os.ReadFile(scriptPath)
	testhelpers.AssertFatalError(t, err)

	_, err = db.Exec(string(script))
	testhelpers.AssertFatalError(t, err)
}

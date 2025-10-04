package database

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:mysecret@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries // Replace with the actual type definition or import the correct package

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	testQueries = New(conn) // Replace with the actual function or import the correct package

	os.Exit(m.Run())
}

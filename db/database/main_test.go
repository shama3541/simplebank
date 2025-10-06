package database

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/shama3541/simplebank/util"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDb *sql.DB

func TestMain(m *testing.M) {
	var err error
	config, err := util.LoadConfig("/Users/shamadeep/workspace/go-workspace/simple_bank/app.env")
	if err != nil {
		log.Println("Error loading config file:", err)
	}
	testDb, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}
	testQueries = New(testDb) // Ensure the `New` function is defined in your code or imported from the correct package

	os.Exit(m.Run())
}

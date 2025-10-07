package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/shama3541/simplebank/api"
	db "github.com/shama3541/simplebank/db/database"
	"github.com/shama3541/simplebank/util"
)

func main() {
	var err error
	config, err := util.LoadConfig("/Users/shamadeep/workspace/go-workspace/simple_bank/app.env")
	if err != nil {
		log.Print("Error loading from config file", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(&store)
	server.StartServer(config.Address)

}

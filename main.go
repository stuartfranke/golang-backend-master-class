package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/stuartfranke/golang-backend-master-class/api"
	db "github.com/stuartfranke/golang-backend-master-class/db/sqlc"
	"github.com/stuartfranke/golang-backend-master-class/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}

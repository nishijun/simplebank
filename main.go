package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/nishijun/simplebank/api"
	db "github.com/nishijun/simplebank/db/sqlc"
	"github.com/nishijun/simplebank/util"
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

	if err := server.Start(config.ServerAddress); err != nil {
		log.Fatal("cannot start server:", err)
	}
}

package main

import (
	"database/sql"
	"log"

	"github.com/gafar-code/online-store/api"
	db "github.com/gafar-code/online-store/db/sqlc"
	"github.com/gafar-code/online-store/util"

	_ "github.com/lib/pq"
)

func main() {
	conf, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Tidak dapat memuat config:", err)
	}

	conn, err := sql.Open(conf.DBDriver, conf.DbSource)
	if err != nil {
		log.Fatal("Tidak bisa terhubung ke database:", err)
	}

	q := db.New(conn)

	server, err := api.NewServer(conf, q)
	if err != nil {
		log.Fatal("Tidak dapat memuat server:", err)
	}

	server.Start(conf.ServerAddress)
}

package main

import (
	"TeslaCoil196/api"
	db "TeslaCoil196/db/sqlc"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgres://root:root@localhost:5432/TeslaBank?sslmode=disable"
	address  = "0.0.0.0:8080"
)

func main() {

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Couldn't connect to databse ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.StartServer(address)
	if err != nil {
		log.Fatal("Unable to start server ", err)
	}
}

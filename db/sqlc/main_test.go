package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgres://root:root@localhost:5432/TeslaBank?sslmode=disable"
)

var testQuries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Couldn't connect to databse ", err)
	}

	testQuries = New(conn)

	os.Exit(m.Run())
}

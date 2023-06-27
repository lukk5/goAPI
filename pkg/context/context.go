package context

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const ( // for now, TODO: move to configs
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "goDB"
)

func NewDbConnection() *sql.DB {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		defer func(db *sql.DB) {
			_ = db.Close()
		}(db)
		log.Fatal(err)
		return nil
	}
	return db
}

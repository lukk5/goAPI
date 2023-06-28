package context

import (
	"fmt"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const ( // for now, TODO: move to configs
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "goDB"
)

func NewDbConnection() *gorm.DB {

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}
	return db
}

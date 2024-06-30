package persistence

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/antonio-lm-jr/products-go/config"
	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	Conn *sql.DB
}

func ProvideDatabase(config config.Config) *Database {
	db := &Database{}
	db.Init(config)

	return db
}

func (Db *Database) Init(config config.Config) {
	var err error

	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.Db.User, config.Db.Pass, config.Db.Host, config.Db.Port, config.Db.Database)

	Db.Conn, err = sql.Open("mysql", connString)

	Db.Conn.SetConnMaxLifetime(time.Minute * 1)
	Db.Conn.SetMaxIdleConns(5)
	Db.Conn.SetMaxOpenConns(5)

	if err != nil {
		log.Println("Error on create Database driver: ", err.Error())
	}

	err = Db.Conn.Ping()

	if err != nil {
		log.Println("Error on open Database: ", err.Error())
	}
}

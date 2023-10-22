package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB

func ConnectDB() {
	db, err := sql.Open("mysql", "root:@/jamijabal?parseTime=true")
	if err != nil {
		panic(err)
	}
	log.Println("Databases Connected")
	DB = db
}

package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB

func ConnectDB() {
	db, err := sql.Open("mysql", "amrizal:hanyasatu@/go_products?parseTime=true")
	if err != nil {
		panic(err)
	}

	log.Println("Database connected")
	DB = db
}

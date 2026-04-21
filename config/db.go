package config

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB(){
	dsn := "root:mahesh123@tcp(127.0.0.1:3306)/blog_news"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	DB = db
}
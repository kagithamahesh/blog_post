package config

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
)

var DB *sql.DB
var RDB *redis.Client
var Ctx = context.Background()

func ConnectDB() {
	dsn := "root:mahesh123@tcp(127.0.0.1:3306)/blog_news"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	DB = db
}

func ConnectRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}

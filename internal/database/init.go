package database

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	client *sql.DB
	once   sync.Once
)

func InitializeDB(dsn string) {
	once.Do(func() {
		var err error
		client, err = sql.Open("mysql", dsn)
		if err != nil {
			log.Fatalf("Error opening database: %v", err)
		}

		if err = client.Ping(); err != nil {
			log.Fatalf("Error pinging database: %v", err)
		} else {
			log.Println("Service connected to Database.")
		}
	})
}

func GetClient() *sql.DB {
	if client == nil {
		log.Fatal("Database has not been initialized. Call database.Initialize first.")
	}
	return client
}

func Close() {
	client.Close()
}

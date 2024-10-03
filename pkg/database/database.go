package database

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// NewDatabase() PostgreSQL initialization
func NewDatabase() (*sqlx.DB, error) {
	err := godotenv.Load() //by default, it is .env so we don't have to write
	if err != nil {
		fmt.Println("Error is occurred  on .env file please check")
	}
	//we read our .env file
	host := os.Getenv("HOST")
	port := os.Getenv("PORT") // don't forget to convert int since port is int type.
	user := os.Getenv("PG_USER")
	dbname := os.Getenv("DB_NAME")
	pass := os.Getenv("PASSWORD")
	dbPgDriver := os.Getenv("DB_PGDRIVER")
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		host,
		port,
		user,
		dbname,
		pass,
	)

	fmt.Println(dataSourceName)

	db, err := sqlx.Connect(dbPgDriver, dataSourceName)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

package postgres

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type DB struct {
	db  *sql.DB
	DSN string
}

type DBAuth struct {
	DB_host     string
	DB_port     string
	DB_user     string
	DB_password string
	DB_name     string
}

func CreateDB() *DB {
	godotenv.Load("../local.env")

	auth := &DBAuth{
		DB_host:     os.Getenv("DB_HOST"),
		DB_port:     os.Getenv("DB_PORT"),
		DB_user:     os.Getenv("DB_USERNAME"),
		DB_password: os.Getenv("DB_PASSWORD"),
		DB_name:     os.Getenv("DB_NAME"),
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", auth.DB_host, auth.DB_port, auth.DB_user, auth.DB_password, auth.DB_name)

	db := &DB{
		DSN: dsn,
	}
	return db
}

// connect to database
func (db *DB) OpenDB() (err error) {
	db.db, err = sql.Open("postgres", db.DSN)
	if err != nil {
		panic(err)
	}

	ping := db.db.Ping()
	if ping != nil {
		panic(ping)
	}

	if db.DSN == "" {
		return fmt.Errorf("dsn is required")
	}

	return nil
}

func (db *DB) InitDB() {
	// insert SQL code here to initialze tables

	query := `
		CREATE TABLE if not exists Users(
			id Serial PRIMARY KEY,
			name VARCHAR(255),
			email VARCHAR(255),
			password VARCHAR(255),
			created_at VARCHAR(255),
			updated_at VARCHAR(255)
		);
	`
	_, err := db.db.Query(query)
	if err != nil {
		panic(err)
	}
}

package repository

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB() (*sql.DB, error) {
	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME")
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	// create db
	_, err = db.Exec(` CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		username VARCHAR(50) NOT NULL,
		password VARCHAR(64) NOT NULL
	)`)

	if err != nil {
		return nil, err
	}

	return db, nil
}

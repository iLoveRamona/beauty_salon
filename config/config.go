package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {
	connStr := "user=pg dbname=beauty_salon_bd password=root host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	fmt.Println("Successfully connected to PostgreSQL!")
	return db, nil
}

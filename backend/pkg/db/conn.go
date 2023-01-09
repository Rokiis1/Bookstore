package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// Conn represents a connection to a PostgreSQL database.
type Conn struct {
	*sql.DB
}

// New creates a new connection to a PostgreSQL database.
func New(host, port, user, password, dbname string) (*Conn, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &Conn{db}, nil
}
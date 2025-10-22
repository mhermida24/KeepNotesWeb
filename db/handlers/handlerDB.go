package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

const dbDriver = "postgres"
const dbSource = "postgres://keepnotes:keepnotes@db:5432/keepnotesdb?sslmode=disable"

func ConnectDB() (*sql.DB, error) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		return nil, err
	}
	if err := conn.Ping(); err != nil {
		return nil, err
	}
	return conn, nil
}

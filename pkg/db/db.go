package db

import (
    "context"
    "log"
    "github.com/jackc/pgx/v5"
)

// GetDBConnection возвращает подключение к базе данных.
func GetDBConnection() *pgx.Conn {
    conn, err := pgx.Connect(context.Background(), "postgres://postgres:your_password@localhost:5432/finance_app")
    if err != nil {
        log.Fatalf("не удалось подключиться к БД: %v", err)
    }
    return conn
}

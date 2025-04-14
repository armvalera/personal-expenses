package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

// DB структура для работы с пулом соединений
type DB struct {
	pool *pgxpool.Pool
}

// NewDB создает новое подключение к базе данных с использованием пула соединений.
func NewDB() (*DB, error) {
	pool, err := pgxpool.New(context.Background(), "postgres://postgres:your_password@localhost:5432/finance_app")
	if err != nil {
		return nil, err
	}
	return &DB{pool: pool}, nil
}

// Close закрывает пул соединений
func (db *DB) Close() {
	db.pool.Close()
}

// GetPool возвращает пул соединений
func (db *DB) GetPool() *pgxpool.Pool {
	return db.pool
}

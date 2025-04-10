package models

// Структура для транзакции
type Transaction struct {
	ID              int
	UserID          int
	CategoryID      int
	Amount          float64
	TransactionDate string
	Description     string
}

// Можно добавить методы для создания, обновления и удаления транзакций.

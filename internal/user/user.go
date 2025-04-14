package user

import (
	"context"
	"fmt"
	"personalexpenses/pkg/db"
	"personalexpenses/pkg/utils"
)

// User структура для представления данных пользователя.
type User struct {
	Username string
	Email    string
	Password string
}

// RegisterUser регистрирует нового пользователя в базе данных.
func RegisterUser(ctx context.Context, user User) error {
	// Хешируем пароль
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("ошибка хеширования пароля: %v", err)
	}

	// Получаем подключение к базе данных через пул
	database, err := db.NewDB()
	if err != nil {
		return fmt.Errorf("не удалось подключиться к базе данных: %v", err)
	}
	defer database.Close()

	// Используем пул соединений
	conn := database.GetPool()

	// Вставляем нового пользователя в таблицу
	_, err = conn.Exec(ctx, "INSERT INTO users (username, email, password) VALUES ($1, $2, $3)", user.Username, user.Email, hashedPassword)
	if err != nil {
		return fmt.Errorf("ошибка при добавлении пользователя в базу данных: %v", err)
	}

	fmt.Printf("Пользователь %s успешно зарегистрирован!\n", user.Username)
	return nil
}

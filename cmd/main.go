package main

import (
	"context"
	"fmt"
	"log"
	"personalexpenses/internal/category"
	"personalexpenses/pkg/db"
)

func main() {
	// Создаем подключение к базе данных
	database, err := db.NewDB()
	if err != nil {
		log.Fatalf("Ошибка при подключении к базе данных: %v", err)
	}
	defer database.Close() // Закрытие пула при завершении работы программы

	fmt.Println("Подключение к базе данных установлено")

	// Контекст для выполнения запросов
	ctx := context.Background()

	// Ввод действия от пользователя
	var action string
	fmt.Println("Выберите действие: create, get, delete")
	fmt.Scanln(&action)

	// Обработка действия
	switch action {
	case "create":
		var categoryName string
		fmt.Println("Введите название категории:")
		fmt.Scanln(&categoryName)
		if err := category.CreateCategory(ctx, database, categoryName); err != nil {
			log.Fatalf("Ошибка при добавлении категории: %v", err)
		}
	case "get":
		categories, err := category.GetCategories(ctx, database)
		if err != nil {
			log.Fatalf("Ошибка при получении категорий: %v", err)
		}
		for _, cat := range categories {
			fmt.Printf("ID: %d, Название: %s\n", cat.ID, cat.Name)
		}
	case "delete":
		var categoryID int
		fmt.Println("Введите ID категории для удаления:")
		fmt.Scanln(&categoryID)
		if err := category.DeleteCategory(ctx, database, categoryID); err != nil {
			log.Fatalf("Ошибка при удалении категории: %v", err)
		}
	default:
		fmt.Println("Неизвестное действие")
	}
}

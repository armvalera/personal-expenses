package main

import (
    "context"
    "fmt"
    "log"
    "personalexpenses/internal/category"
)

func main() {
    var action string
    fmt.Println("Выберите действие: create, get, delete")
    fmt.Scanln(&action)

    ctx := context.Background()

    switch action {
    case "create":
        var categoryName string
        fmt.Println("Введите название категории:")
        fmt.Scanln(&categoryName)
        if err := category.CreateCategory(ctx, categoryName); err != nil {
            log.Fatalf("Ошибка при добавлении категории: %v", err)
        }
    case "get":
        categories, err := category.GetCategories(ctx)
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
        if err := category.DeleteCategory(ctx, categoryID); err != nil {
            log.Fatalf("Ошибка при удалении категории: %v", err)
        }
    default:
        fmt.Println("Неизвестное действие")
    }
}

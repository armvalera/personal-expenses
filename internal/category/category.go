package category

import (
    "context"
    "fmt"
    "personalexpenses/pkg/db"
)

// Category структура для представления категории.
type Category struct {
    ID   int
    Name string
}

// CreateCategory создает новую категорию в базе данных.
func CreateCategory(ctx context.Context, name string) error {
    conn := db.GetDBConnection()
    defer conn.Close(ctx)

    // Вставляем категорию в таблицу
    _, err := conn.Exec(ctx, "INSERT INTO categories (name) VALUES ($1)", name)
    if err != nil {
        return fmt.Errorf("ошибка при добавлении категории в базу данных: %v", err)
    }

    fmt.Printf("Категория '%s' успешно добавлена!\n", name)
    return nil
}

// GetCategories получает все категории из базы данных.
func GetCategories(ctx context.Context) ([]Category, error) {
    conn := db.GetDBConnection()
    defer conn.Close(ctx)

    rows, err := conn.Query(ctx, "SELECT id, name FROM categories")
    if err != nil {
        return nil, fmt.Errorf("ошибка при получении категорий из базы данных: %v", err)
    }
    defer rows.Close()

    var categories []Category
    for rows.Next() {
        var category Category
        if err := rows.Scan(&category.ID, &category.Name); err != nil {
            return nil, fmt.Errorf("ошибка при сканировании категории: %v", err)
        }
        categories = append(categories, category)
    }

    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("ошибка при обходе строк: %v", err)
    }

    return categories, nil
}

// DeleteCategory удаляет категорию по ID.
func DeleteCategory(ctx context.Context, categoryID int) error {
    conn := db.GetDBConnection()
    defer conn.Close(ctx)

    // Удаляем категорию из базы данных
    _, err := conn.Exec(ctx, "DELETE FROM categories WHERE id = $1", categoryID)
    if err != nil {
        return fmt.Errorf("ошибка при удалении категории: %v", err)
    }

    fmt.Printf("Категория с ID %d успешно удалена!\n", categoryID)
    return nil
}

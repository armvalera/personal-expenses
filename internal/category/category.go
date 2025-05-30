package category

import (
    "context"
    "fmt"
    "personalexpenses/pkg/db"
)

type Category struct {
    ID   int
    Name string
}

func CreateCategory(ctx context.Context, db *db.DB, name string) error {
    conn := db.GetPool()
    _, err := conn.Exec(ctx, "INSERT INTO categories (name) VALUES ($1)", name)
    if err != nil {
        return fmt.Errorf("ошибка при добавлении категории в базу данных: %v", err)
    }
    fmt.Printf("Категория '%s' успешно добавлена!\n", name)
    return nil
}

func GetCategories(ctx context.Context, db *db.DB) ([]Category, error) {
    conn := db.GetPool()
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

func DeleteCategory(ctx context.Context, db *db.DB, categoryID int) error {
    conn := db.GetPool()
    _, err := conn.Exec(ctx, "DELETE FROM categories WHERE id = $1", categoryID)
    if err != nil {
        return fmt.Errorf("ошибка при удалении категории: %v", err)
    }
    fmt.Printf("Категория с ID %d успешно удалена!\n", categoryID)
    return nil
}

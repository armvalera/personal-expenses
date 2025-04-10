package utils

import (
    "golang.org/x/crypto/bcrypt"
    "fmt"
)

// HashPassword хеширует пароль с использованием bcrypt.
func HashPassword(password string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", fmt.Errorf("не удалось хешировать пароль: %v", err)
    }
    return string(hashedPassword), nil
}

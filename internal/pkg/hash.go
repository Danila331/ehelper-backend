package pkg

import (
	"golang.org/x/crypto/bcrypt"
)

// Функция для создания хэша пароля
func HashPassword(password string) (string, error) {
	// Генерация хэша для пароля
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// Функция для проверки пароля
func CheckPassword(inputPassword, hashedPassword string) bool {
	// Сравниваем введённый пароль с закешированным паролем
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
	return err == nil
}

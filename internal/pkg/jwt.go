package pkg

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	// Секретный ключ для подписи токена
	secretKey = []byte("I love mami")
)

// Claims содержит информацию, которую вы хотите закодировать в токене
type Claims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

// Генерация токена
func GenerateToken(userID int) (string, error) {
	// Создание структуры с данными для токена
	claims := Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 160).Unix(), // Токен действителен в течение 24 часов
		},
	}

	// Создание токена с помощью структуры claims и секретного ключа
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Подписываем токен и возвращаем его в виде строки
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Функция для извлечения ID пользователя из токена
func ExtractUserIDFromToken(tokenString string) (int, error) {
	// Парсинг токена с помощью секретного ключа
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return 0, err
	}

	// Проверка валидности токена
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.UserID, nil
	} else {
		return 0, fmt.Errorf("invalid token")
	}
}

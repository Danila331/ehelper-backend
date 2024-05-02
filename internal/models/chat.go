package models

import (
	"fmt"

	"github.com/Danila331/mifiotsos/internal/store"
)

type Chat struct {
	Id       int
	UserId   string
	UserName string
	ChatId   string
	Anger    int
	Disgust  int
	Fear     int
	Happy    int
	Neutral  int
	Sad      int
	Suprised int
}

type ChatAverage struct {
	Id       int
	UserId   string
	UserName string
	ChatId   string
	Anger    float64
	Disgust  float64
	Fear     float64
	Happy    float64
	Neutral  float64
	Sad      float64
	Suprised float64
}
type ChatInterface interface {
	ReadAll() ([]Chat, error)
	ReadAllByAvr() ([]ChatAverage, error)
}

func (c *Chat) ReadAllByAvr(chatid string) ([]ChatAverage, error) {
	db, err := store.ConnectDB()
	if err != nil {
		return []ChatAverage{}, err
	}
	defer db.Close()
	rows, err := db.Query(`
		SELECT "username",
		       ROUND(AVG("anger"),1) AS "avg_anger",
		       ROUND(AVG("disgust"),1) AS "avg_disgust",
		       ROUND(AVG("fear"),1) AS "avg_fear",
		       ROUND(AVG("happy"),1) AS "avg_happy",
		       ROUND(AVG("neutral"),1) AS "avg_neutral",
		       ROUND(AVG("sad"),1) AS "avg_sad",
		       ROUND(AVG("surprised"),1) AS "avg_suprised"
		FROM "msgs"
		WHERE chatid=$1
		GROUP BY "username"
	`, chatid)
	if err != nil {
		return []ChatAverage{}, err
	}

	defer rows.Close()

	var chats []ChatAverage
	for rows.Next() {
		var chat ChatAverage
		err = rows.Scan(&chat.UserName, &chat.Anger, &chat.Disgust, &chat.Fear, &chat.Happy, &chat.Neutral, &chat.Sad, &chat.Suprised)
		if err != nil {
			return []ChatAverage{}, err
		}
		chats = append(chats, chat)
	}

	return chats, nil

}
func (c *Chat) ReadAll() ([]Chat, error) {
	db, err := store.ConnectDB()
	if err != nil {
		return []Chat{}, err
	}
	defer db.Close()
	rows, err := db.Query(`SELECT * FROM msgs`)
	if err != nil {
		return nil, fmt.Errorf("ошибка при выполнении запроса SELECT: %v", err)
	}
	defer rows.Close()

	var chats []Chat

	// Итерация по результатам запроса
	for rows.Next() {
		var chat Chat
		// Сканирование данных из строки результата в переменные структуры Conference
		err := rows.Scan(&chat.Id, &chat.UserId, &chat.UserName, &chat.ChatId, &chat.Anger, &chat.Disgust, &chat.Fear, &chat.Happy, &chat.Neutral, &chat.Sad, &chat.Suprised)
		if err != nil {
			return nil, fmt.Errorf("ошибка при сканировании строки результата: %v", err)
		}
		// Добавление конференции в срез
		chats = append(chats, chat)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка при итерации по результатам запроса: %v", err)
	}

	return chats, nil
}

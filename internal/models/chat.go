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

type ChatInterface interface {
	ReadAll() ([]Chat, error)
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

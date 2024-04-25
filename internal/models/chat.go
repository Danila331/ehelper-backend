package models

import (
	"fmt"

	"github.com/Danila331/mifiotsos/internal/store"
)

type Chat struct {
	Id          int
	UserId      string
	UserName    string
	ChatId      string
	Calm        int
	Disgust     int
	Openness    int
	Sociability int
	Anger       int
	Balance     int
	Depression  int
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
		err := rows.Scan(&chat.Id, &chat.UserId, &chat.UserName, &chat.ChatId, &chat.Calm, &chat.Disgust, &chat.Openness, &chat.Sociability, &chat.Anger, &chat.Balance, &chat.Depression)
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

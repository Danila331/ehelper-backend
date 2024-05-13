package models

import (
	"fmt"

	"github.com/Danila331/mifiotsos/internal/store"
)

// Структура сообщений
type Msg struct {
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

// Структура сообщений только со среднем значением
type MsgAverage struct {
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

// Интерфейс для реализации методов
type MsgInterface interface {
	ReadAll() ([]Msg, error)
	ReadAllByAvr() ([]MsgAverage, error)
}

// Метод для получения среднего по пользователям и чатам
func (m *Msg) ReadAllByAvr(chatsid string) ([]MsgAverage, error) {
	db, err := store.ConnectDB()
	if err != nil {
		return []MsgAverage{}, err
	}
	defer db.Close()
	chatsidString := fmt.Sprintf("%s%s", "'000000', '000000' ", chatsid)
	rows, err := db.Query(fmt.Sprintf(`
		SELECT "username",
		       ROUND(AVG("anger"),1) AS "avg_anger",
		       ROUND(AVG("disgust"),1) AS "avg_disgust",
		       ROUND(AVG("fear"),1) AS "avg_fear",
		       ROUND(AVG("happy"),1) AS "avg_happy",
		       ROUND(AVG("neutral"),1) AS "avg_neutral",
		       ROUND(AVG("sad"),1) AS "avg_sad",
		       ROUND(AVG("surprised"),1) AS "avg_suprised"
		FROM "msgs"
		WHERE chatid IN (%s)
		GROUP BY "username"
	`, chatsidString))
	if err != nil {
		return []MsgAverage{}, err
	}

	defer rows.Close()

	var msgs []MsgAverage
	for rows.Next() {
		var msg MsgAverage
		err = rows.Scan(&msg.UserName, &msg.Anger, &msg.Disgust, &msg.Fear, &msg.Happy, &msg.Neutral, &msg.Sad, &msg.Suprised)
		if err != nil {
			return []MsgAverage{}, err
		}
		msgs = append(msgs, msg)
	}

	return msgs, nil
}

// Метод для получения всех сообщений из базыданных
func (m *Msg) ReadAll(chatsid string) ([]Msg, error) {
	db, err := store.ConnectDB()
	if err != nil {
		return []Msg{}, err
	}
	defer db.Close()
	chatsidString := fmt.Sprintf("%s%s", "'000000', '000000' ", chatsid)
	rows, err := db.Query(fmt.Sprintf(`SELECT * FROM msgs WHERE chatid IN (%s)`, chatsidString))
	if err != nil {
		return nil, fmt.Errorf("ошибка при выполнении запроса SELECT: %v", err)
	}
	defer rows.Close()

	var msgs []Msg

	// Итерация по результатам запроса
	for rows.Next() {
		var msg Msg
		// Сканирование данных из строки результата в переменные структуры Conference
		err := rows.Scan(&msg.Id, &msg.UserId, &msg.UserName, &msg.ChatId, &msg.Anger, &msg.Disgust, &msg.Fear, &msg.Happy, &msg.Neutral, &msg.Sad, &msg.Suprised)
		if err != nil {
			return nil, fmt.Errorf("ошибка при сканировании строки результата: %v", err)
		}
		// Добавление конференции в срез
		msgs = append(msgs, msg)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка при итерации по результатам запроса: %v", err)
	}

	return msgs, nil
}

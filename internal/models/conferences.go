package models

import (
	"fmt"
	"time"

	"github.com/Danila331/mifiotsos/internal/store"
)

type Conferences struct {
	Id         int
	Date       time.Time
	Status     string
	Filepath   string
	Anger      float64
	Disgust    float64
	Enthusiasm float64
	Fear       float64
	Happiness  float64
	Neutral    float64
	Sadness    float64
}

type ConferencesInterface interface {
	Create() error
	ReadAll() ([]Conferences, error)
}

func (c *Conferences) Create() error {
	db, err := store.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()
	query := `
        INSERT INTO conferences (date, filepath, status, anger, disgust, enthusiasm, fear, happiness, neutral, sadness) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
    `
	DataTime := time.Now()
	_, err = db.Exec(query, DataTime, c.Filepath, "no", 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05)
	if err != nil {
		return fmt.Errorf("ошибка при вставке данных: %v", err)
	}

	return nil
}

func (c *Conferences) ReadAll() ([]Conferences, error) {
	db, err := store.ConnectDB()
	if err != nil {
		return []Conferences{}, err
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM conferences")
	if err != nil {
		return nil, fmt.Errorf("ошибка при выполнении запроса SELECT: %v", err)
	}
	defer rows.Close()

	var conferences []Conferences

	// Итерация по результатам запроса
	for rows.Next() {
		var conference Conferences
		// Сканирование данных из строки результата в переменные структуры Conference
		err := rows.Scan(&conference.Id, &conference.Date, &conference.Status, &conference.Filepath, &conference.Anger, &conference.Disgust, &conference.Enthusiasm, &conference.Fear, &conference.Happiness, &conference.Neutral, &conference.Sadness)
		conference.Anger = conference.Anger * 100
		conference.Disgust = conference.Disgust * 100
		conference.Enthusiasm = conference.Enthusiasm * 100
		conference.Fear = conference.Fear * 100
		conference.Happiness = conference.Happiness * 100
		conference.Neutral = conference.Neutral * 100
		conference.Sadness = conference.Sadness * 100
		if err != nil {
			return nil, fmt.Errorf("ошибка при сканировании строки результата: %v", err)
		}
		// Добавление конференции в срез
		conferences = append(conferences, conference)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка при итерации по результатам запроса: %v", err)
	}

	return conferences, nil
}

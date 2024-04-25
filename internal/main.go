package main

import (
	"fmt"

	"github.com/Danila331/mifiotsos/internal/server"
	"github.com/Danila331/mifiotsos/internal/store"
)

func main() {
	db, err := store.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	if _, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            Id SERIAL PRIMARY KEY,
            Password TEXT NOT NULL,
            Email TEXT UNIQUE NOT NULL,
            ChatsId TEXT,
            Status TEXT
        );
    `); err != nil {
		panic(err)
	}
	// Создание таблицы конференций
	_, err = db.Exec(`
    CREATE TABLE IF NOT EXISTS conferences (
        id SERIAL PRIMARY KEY,
        date DATE NOT NULL,
		filepath TEXT,
		status TEXT,
        anger FLOAT,
        disgust FLOAT,
        enthusiasm FLOAT,
        fear FLOAT,
        happiness FLOAT,
        neutral FLOAT,
        sadness FLOAT
		)`)

	if err != nil {
		panic(err)
	}

	fmt.Println("Таблица конференций успешно создана")
	server.StartServer()
}

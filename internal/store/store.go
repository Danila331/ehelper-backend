package store

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// POSTGRESQL_HOST=147.45.144.139
// POSTGRESQL_PORT=5432
// POSTGRESQL_USER=ehelperbot
// POSTGRESQL_PASSWORD=FM)I?EDJs&b1?M
// POSTGRESQL_DBNAME=ehelper

func ConnectDB() (*sql.DB, error) {
	// Строка подключения к базе данных PostgreSQL
	connectionString := "postgres://ehelperbot:FM%29I%3FEDJs%26b1%3FM@147.45.144.139:5432/ehelper"

	// Открываем соединение с базой данных
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	// Проверяем, что соединение установлено успешно
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	// Возвращаем объект соединения
	return db, nil
}

package models

import (
	"fmt"

	"github.com/Danila331/mifiotsos/internal/store"
)

// Структура user
type User struct {
	Id       int    //Id
	Password string //Password
	Email    string //Email
	ChatsId  string //ChatsId id чатов ввиде строки разделены ;
	Status   string //Status для подписки, тип подписки
}

// Интерфейс который реализует все методы user
type UserINterface interface {
	Create() error
	Update() error
	ReadByEmail() (User, error)
}

// Метод для создавания user в бд
func (u *User) Create() error {
	conn, err := store.ConnectDB()
	if err != nil {
		return err
	}
	defer conn.Close()
	query := `INSERT INTO "users" (password, email)
              VALUES ($1, $2)`

	// Выполнение SQL-запроса
	_, err = conn.Exec(query, u.Password, u.Email)
	if err != nil {
		return fmt.Errorf("failed to insert user: %v", err)
	}
	return nil
}

// Метод для получения пользователя по почте
func (u *User) ReadByEmail(email string) (User, error) {
	conn, err := store.ConnectDB()
	if err != nil {
		return User{}, err
	}
	defer conn.Close()

	query := "SELECT id, password, email, chatsid FROM users WHERE email=$1"
	row := conn.QueryRow(query, email)

	var user User
	err = row.Scan(&user.Id, &user.Password, &user.Email, &user.ChatsId)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// Метод для обновления данных о пользователе
func (u *User) Update() error {
	conn, err := store.ConnectDB()
	if err != nil {
		return err
	}
	defer conn.Close()

	query := `UPDATE "users" SET chatsid=$2 WHERE email=$1`

	// Выполнение SQL-запроса
	_, err = conn.Exec(query, u.Email, u.ChatsId)
	if err != nil {
		return fmt.Errorf("failed to update user: %v", err)
	}

	return nil
}

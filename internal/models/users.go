package models

import (
	"fmt"

	"github.com/Danila331/mifiotsos/internal/store"
)

type User struct {
	Id       int    //Id
	Password string //Password
	Email    string //Email
	ChatsId  string //ChatsId id чатов ввиде строки разделены ;
	Status   string //Status для подписки, тип подписки
}

type UserINterface interface {
	Create() error
	Update() error
	ReadByEmail() (User, error)
}

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

func (u *User) ReadByEmail(email string) (User, error) {
	conn, err := store.ConnectDB()
	if err != nil {
		return User{}, err
	}
	defer conn.Close()

	query := "SELECT id, password, email FROM users WHERE email=$1"
	row := conn.QueryRow(query, email)

	var user User
	err = row.Scan(&user.Id, &user.Password, &user.Email)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (u *User) Update() error {
	conn, err := store.ConnectDB()
	if err != nil {
		return err
	}
	defer conn.Close()

	query := `UPDATE "user"
              SET email=$2, password=$3, chats_id=$4, status=$5
              WHERE id=$1`

	// Выполнение SQL-запроса
	_, err = conn.Exec(query, u.Id, u.Password, u.Email, u.ChatsId, u.Status)
	if err != nil {
		return fmt.Errorf("failed to update user: %v", err)
	}

	return nil
}

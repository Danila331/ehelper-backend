package forms

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"time"

	"github.com/Danila331/mifiotsos/internal/models"
	"github.com/Danila331/mifiotsos/internal/pkg"
	"github.com/labstack/echo/v4"
)

// Функция для обработки формы регистрации
func SignForm(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	hashPassword, err := pkg.HashPassword(password)
	if err != nil {
		return err
	}
	user := models.User{
		Password: hashPassword,
		Email:    email,
	}
	err = user.Create()
	if err != nil {
		errorWeb := models.ErrorWeb{Number: "409", ErrorString: "Пользователь с такой почтой уже есть, попробуйте войти.", BackLinkText: "Войти", BackLink: "login"}
		_ = errorWeb.CreatePage(c)
		return err
	}
	tokenString, err := pkg.GenerateToken(email, hashPassword)
	if err != nil {
		return err
	}

	cookie := http.Cookie{
		Name:    "jwt",
		Value:   tokenString,
		Expires: time.Now().Add(time.Hour * 640),
		Path:    "/",
	}

	http.SetCookie(c.Response(), &cookie)
	fmt.Println(cookie)

	htmlFiles := []string{
		filepath.Join("./", "templates", "submit", "sign_submit.html"),
	}

	templ, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		return err
	}

	templ.ExecuteTemplate(c.Response(), "sign_submit", nil)
	return nil
}

package forms

import (
	"net/http"
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

	// Проверяем есть ли такой пользователь с указанной почтой
	if err != nil {
		errorWeb := models.ErrorWeb{Number: "409", ErrorString: "Пользователь с такой почтой уже есть, попробуйте войти.", BackLinkText: "Войти", BackLink: "login"}
		_ = errorWeb.CreatePage(c)
		return err
	}

	// Создаем JWT токен для пользователя
	tokenString, err := pkg.GenerateToken(email, hashPassword)
	if err != nil {
		return err
	}

	// Кладем JWT в Cookie
	cookie := http.Cookie{
		Name:    "jwt",
		Value:   tokenString,
		Expires: time.Now().Add(time.Hour * 640),
		Path:    "/",
	}

	http.SetCookie(c.Response(), &cookie)

	err = pkg.HtmlPageRender("submit/sign_submit.html", "sign_submit", c)

	if err != nil {
		return err
	}

	return nil
}

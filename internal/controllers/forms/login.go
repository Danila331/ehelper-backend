package forms

import (
	"net/http"
	"time"

	"github.com/Danila331/mifiotsos/internal/models"
	"github.com/Danila331/mifiotsos/internal/pkg"
	"github.com/labstack/echo/v4"
)

// Функция для обработки формы Login
func LoginForm(c echo.Context) error {
	password := c.FormValue("password")
	email := c.FormValue("email")
	var user models.User
	user, err := user.ReadByEmail(email)

	// Ошибка если пользователь с такой почтой не найден
	if err != nil {
		errorWeb := models.ErrorWeb{Number: "404", ErrorString: "Такого пользователя не существует, зарегистрируйтесь.", BackLinkText: "Регистрация", BackLink: "sign"}
		_ = errorWeb.CreatePage(c)
		return err
	}

	// Ошибка если пользователь указал неверный пароль
	if !pkg.CheckPassword(password, user.Password) {
		errorWeb := models.ErrorWeb{Number: "535", ErrorString: "Неверный пароль, попробуйте еще раз", BackLinkText: "Назад", BackLink: "login"}
		_ = errorWeb.CreatePage(c)
		return err
	}

	//Создаем JWT токен для пользователя
	tokenString, err := pkg.GenerateToken(email, password)
	if err != nil {
		return err
	}

	//Кладем JWT в Cookie
	cookie := http.Cookie{
		Name:    "jwt",
		Value:   tokenString,
		Expires: time.Now().Add(time.Hour * 640),
		Path:    "/",
	}

	http.SetCookie(c.Response(), &cookie)

	err = pkg.HtmlPageRender("submit/login_submit.html", "login_submit", c)
	if err != nil {
		return err
	}
	return nil
}

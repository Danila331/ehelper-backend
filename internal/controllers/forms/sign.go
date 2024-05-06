package forms

import (
	"html/template"
	"path/filepath"

	"github.com/Danila331/mifiotsos/internal/models"
	"github.com/labstack/echo/v4"
)

// Функция для обработки формы регистрации
func SignForm(c echo.Context) error {
	user := models.User{
		Password: c.FormValue("password"),
		Email:    c.FormValue("email"),
	}
	err := user.Create()
	if err != nil {
		errorWeb := models.ErrorWeb{Number: "409", ErrorString: "Пользователь с такой почтой уже есть, попробуйте войти.", BackLinkText: "Войти", BackLink: "login"}
		_ = errorWeb.CreatePage(c)
		return nil
	}

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

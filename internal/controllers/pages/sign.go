package pages

import (
	"path/filepath"
	"text/template"

	"github.com/labstack/echo/v4"
)

// Функция для отобрадения страницы регистрации
func SignPage(c echo.Context) error {
	htmlFiles := []string{
		filepath.Join("./", "templates", "forms", "sign.html"),
	}

	templ, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		return err
	}

	templ.ExecuteTemplate(c.Response(), "sign", nil)
	return nil
}

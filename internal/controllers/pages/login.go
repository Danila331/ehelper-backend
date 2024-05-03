package pages

import (
	"path/filepath"
	"text/template"

	"github.com/labstack/echo/v4"
)

// Функция для отобрадения страницы login
func LoginPage(c echo.Context) error {
	htmlFiles := []string{
		filepath.Join("./", "templates", "forms", "login.html"),
	}

	templ, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		return err
	}

	templ.ExecuteTemplate(c.Response(), "login", nil)
	return nil
}

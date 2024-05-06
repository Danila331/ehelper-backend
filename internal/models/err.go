package models

import (
	"html/template"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

// Структура для отображения ошибко на сайте
type ErrorWeb struct {
	Number       string
	ErrorString  string
	BackLinkText string
	BackLink     string
}

type ErrorWebInterface interface {
	CreatePage(echo.Context) error
}

func (e *ErrorWeb) CreatePage(c echo.Context) error {
	htmlFiles := []string{
		filepath.Join("./", "templates", "submit", "err.html"),
	}

	templ, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		return err
	}

	templ.ExecuteTemplate(c.Response(), "err", e)
	return nil
}

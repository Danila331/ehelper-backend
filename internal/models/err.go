package models

// Структура для отображения ошибко на сайте
type ErrorWeb struct {
	Number       string
	ErrorString  string
	BackLinkText string
	BackLink     string
}

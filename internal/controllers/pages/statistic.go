package pages

import (
	"fmt"
	"path/filepath"
	"text/template"

	"github.com/Danila331/mifiotsos/internal/models"
	"github.com/labstack/echo/v4"
)

// Функция для отображения статистики конференций
func StatisticPageConf(c echo.Context) error {
	var conference models.Conferences
	conferences, err := conference.ReadAll()
	fmt.Println(err)
	if err != nil {
		return err
	}

	htmlFiles := []string{
		filepath.Join("./", "templates", "statisticconf.html"),
	}

	templ, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		return err
	}

	templ.ExecuteTemplate(c.Response(), "statistic", conferences)
	return nil
}

// Функция для отобрадения статистики чатов
func StatisticPageChat(c echo.Context) error {
	var msg models.Msg
	var user models.User

	user, err := user.ReadByEmail(c.Get("email").(string))
	if err != nil {
		return err
	}
	msgs, err := msg.ReadAllByAvr(user.ChatsId)
	fmt.Println(err)
	if err != nil {
		return err
	}

	htmlFiles := []string{
		filepath.Join("./", "templates", "statisticchat.html"),
	}

	templ, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		return err
	}

	templ.ExecuteTemplate(c.Response(), "statistic", msgs)
	return nil
}

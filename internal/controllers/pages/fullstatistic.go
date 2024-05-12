package pages

import (
	"fmt"
	"path/filepath"
	"text/template"

	"github.com/Danila331/mifiotsos/internal/models"
	"github.com/Danila331/mifiotsos/internal/pkg"
	"github.com/labstack/echo/v4"
)

// Функция для отображения полной статистики по конференциям
func FulStatisticPageConf(c echo.Context) error {

	fullstatistic, err := pkg.GetFulResultConf()
	if err != nil {
		return err
	}

	htmlFiles := []string{
		filepath.Join("./", "templates", "fulstatisticconf.html"),
	}

	templ, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		return err
	}

	templ.ExecuteTemplate(c.Response(), "fulstatistic", fullstatistic)
	return nil
}

// Функция для отображения полной статистики по чатам
func FulStatisticPageChat(c echo.Context) error {
	var user models.User
	user, err := user.ReadByEmail(c.Get("email").(string))
	if err != nil {
		return err
	}

	fullstatistic, err := pkg.GetFulResultChat(user.ChatsId)
	fmt.Println(fullstatistic)
	if err != nil {
		return err
	}

	htmlFiles := []string{
		filepath.Join("./", "templates", "fulstatisticchat.html"),
	}

	templ, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		return err
	}

	templ.ExecuteTemplate(c.Response(), "fulstatisticchat", fullstatistic)
	return nil
}

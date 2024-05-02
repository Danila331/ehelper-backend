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

func LoginForm(c echo.Context) error {
	password := c.FormValue("password")
	email := c.FormValue("email")
	var user models.User
	user, err := user.ReadByEmail(email)

	if err != nil {
		return err
	}

	if password != user.Password {
		htmlFiles := []string{
			filepath.Join("./", "templates", "submit", "err.html"),
		}

		templ, err := template.ParseFiles(htmlFiles...)
		if err != nil {
			return err
		}

		errorWeb := models.ErrorWeb{Number: "535", ErrorWeb: "Неверный пароль, попробуйте еще раз", BackLink: "login"}

		templ.ExecuteTemplate(c.Response(), "err", errorWeb)
		return nil
	}

	tokenString, err := pkg.GenerateToken(email, password)
	if err != nil {
		return err
	}

	cookie := new(http.Cookie)
	// Name:     "jwt",
	// Value:    tokenString,
	// Expires:  time.Now().Add(time.Hour * 160), // Пример: установка срока действия куки на 24 часа
	// HttpOnly: true,
	// Secure:   false, // Установите true для HTTPS
	// SameSite: http.SameSiteStrictMode,
	cookie.Name = "jwt"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(time.Hour * 160)
	cookie.HttpOnly = true
	cookie.Secure = false
	cookie.SameSite = http.SameSiteNoneMode

	c.SetCookie(cookie)
	fmt.Println(cookie)

	htmlFiles := []string{
		filepath.Join("./", "templates", "submit", "login_submit.html"),
	}

	templ, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		return err
	}

	templ.ExecuteTemplate(c.Response(), "login_submit", nil)
	return nil
}

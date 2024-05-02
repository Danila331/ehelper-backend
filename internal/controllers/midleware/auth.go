package midleware

import (
	"fmt"
	"net/http"

	"github.com/Danila331/mifiotsos/internal/pkg"
	"github.com/labstack/echo/v4"
)

// Middleware, который извлекает JWT из куки и проверяет его подлинность
func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Получаем куки из запроса
		cookie, err := c.Cookie("jwt")
		fmt.Println(cookie, err)
		if err != nil {
			// Если куки отсутствуют, возвращаем ошибку
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}

		tokenString := cookie.Value

		email, err := pkg.ExtractEmailFromToken(tokenString)

		if err != nil {
			return err
		}

		c.Set("email", email)

		// Продолжаем выполнение следующего обработчика
		return next(c)
	}
}

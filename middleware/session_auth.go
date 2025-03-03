package middleware

import (
	"dedi_crm/config"
	"dedi_crm/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SessionMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		// pengecekkan session dan username
		session, _ := config.Store.Get(c.Request(), "session")
		email, ok := session.Values["email"].(string)

		// jika tidak ditemukan, maka dikembalikan pesan error
		if !ok || email == "" {
			utils.Logger.Error("Unauthorized")
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"code":    401,
				"message": "Unauthorized",
			})
		}

		// jika berhasil, maka diteruskan dengan menyimpan data username yang dapat dipanggil dari mana saja
		c.Set("email", email)
		return next(c)
	}
}

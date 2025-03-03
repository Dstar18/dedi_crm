package middleware

import (
	"dedi_crm/config"
	"dedi_crm/models"
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

func IsLogin(valueRole string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			emailIsLogin := c.Get("email").(string)

			var userM models.User
			// Mencari user berdasarkan email
			if err := config.DB.Where("email = ?", emailIsLogin).First(&userM).Error; err != nil {
				utils.Logger.Error("User not found")
				return c.JSON(http.StatusNotFound, map[string]interface{}{
					"code":    http.StatusNotFound,
					"message": "User not found",
				})
			}

			// Memeriksa role pengguna dengan role yang diperlukan
			if userM.Role != valueRole {
				utils.Logger.Error("Forbidden: insufficient role")
				return c.JSON(http.StatusForbidden, map[string]interface{}{
					"code":    http.StatusForbidden,
					"message": "Forbidden: insufficient role",
				})
			}
			return next(c)
		}
	}
}

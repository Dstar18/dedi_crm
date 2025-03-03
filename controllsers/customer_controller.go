package controllsers

import (
	"dedi_crm/config"
	"dedi_crm/models"
	"dedi_crm/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Customers(c echo.Context) error {
	var customerM []models.Lead
	config.DB.Find(&customerM)

	// return success
	utils.Logger.Info("Customers successfully")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"message": "Customers successfully",
		"data":    customerM,
	})
}

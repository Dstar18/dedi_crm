package controllsers

import (
	"dedi_crm/config"
	"dedi_crm/models"
	"dedi_crm/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Products(c echo.Context) error {
	var productM []models.Product
	config.DB.Find(&productM)

	// return success
	utils.Logger.Info("Product successfully")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Product successfully",
		"data":    productM,
	})
}

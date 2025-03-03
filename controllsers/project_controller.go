package controllsers

import (
	"dedi_crm/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ProjectAdd(c echo.Context) error {
	// return success
	utils.Logger.Info("Project Created successfully")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Project Created successfully",
		"Data":    nil,
	})
}

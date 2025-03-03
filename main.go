package main

import (
	"dedi_crm/config"
	"dedi_crm/controllsers"
	"dedi_crm/models"
	"dedi_crm/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {

	// Initialize logger
	utils.InitLogger()

	// connect database
	config.InitDB()

	// Migrate
	config.DB.AutoMigrate(&models.User{})
	config.DB.AutoMigrate(&models.Lead{})
	config.DB.AutoMigrate(&models.Product{})
	config.DB.AutoMigrate(&models.Project{})

	// Seeder
	config.Seed()

	// initialize echo
	e := echo.New()

	// Route public
	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/register", controllsers.StoreUser)

	e.Logger.Fatal(e.Start(":3000"))
}

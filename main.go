package main

import (
	"dedi_crm/config"
	"dedi_crm/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {

	// connect database
	config.InitDB()

	// Migrate
	config.DB.AutoMigrate(&models.User{})
	config.DB.AutoMigrate(&models.Lead{})
	config.DB.AutoMigrate(&models.Product{})

	// Seeder
	config.Seed()

	// initialize echo
	e := echo.New()

	// Route public
	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":3000"))
}

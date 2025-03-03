package main

import (
	"dedi_crm/config"
	"dedi_crm/controllsers"
	"dedi_crm/middleware"
	"dedi_crm/models"
	"dedi_crm/utils"

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

	e.POST("/login", controllsers.Login)
	e.POST("/register", controllsers.StoreUser)
	e.GET("/logout", controllsers.Logout)

	// Route Auth
	// Middleware Session
	protected := e.Group("/api")
	protected.Use(middleware.SessionMiddleware)

	protected.GET("/customers", controllsers.Customers)
	protected.POST("/customers/add", controllsers.CustomerStore)
	protected.GET("/customers/lead", controllsers.CustomerLead)

	protected.GET("/product", controllsers.Products)
	protected.POST("/product/add", controllsers.ProductStore)
	protected.POST("/product/update/:id", controllsers.ProductUpdate)

	e.Logger.Fatal(e.Start(":3000"))
}

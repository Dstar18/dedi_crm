package config

import (
	"dedi_crm/models"
	"dedi_crm/utils"
	"time"
)

// seeder
func Seed() {

	// seed user
	var userM []models.User
	DB.Find(&userM)
	hashedPassword, err := utils.HashPassword("Admin123#")
	if err != nil {
		utils.Logger.Error("Failed to hash password")
	}
	if len(userM) == 0 {
		paramUser := []models.User{
			{Name: "Admin DWP", Email: "admin@mail.com", Password: hashedPassword, Role: "admin", CreatedAt: time.Now()},
		}
		DB.Create(&paramUser)
	}
}

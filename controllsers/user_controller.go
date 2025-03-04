package controllsers

import (
	"dedi_crm/config"
	"dedi_crm/models"
	"dedi_crm/utils"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type UserValidate struct {
	Name     string `json:"name" validate:"required,min=2,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=20"`
	Role     string `json:"role" validate:"required,min=2,max=20"`
}

func StoreUser(c echo.Context) error {
	// request struct validation
	var user UserValidate

	// request params, and check body
	if err := c.Bind(&user); err != nil {
		utils.Logger.Error("Invalid request body")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"message": "invalid request body",
		})
	}

	// validation struc
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = "This field is" + " " + err.Tag() + " " + err.Param()
		}
		utils.Logger.Error(errors)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"message": errors,
		})
	}

	// validation password
	if err := utils.ValidatePassword(user.Password); err != nil {
		utils.Logger.Warn(err.Error())
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"message": err.Error(),
		})
	}

	// request struct model
	var userM models.User

	// check username is ready
	if err := config.DB.Where("email = ? ", user.Email).First(&userM).Error; err == nil {
		utils.Logger.Warn("email " + user.Email + " is already")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":   400,
			"mesage": "email " + user.Email + " is already",
		})
	}

	// hash password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		utils.Logger.Error("Failed to hash password")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    500,
			"message": "Failed to hash password",
		})
	}

	param := models.User{
		Name:      user.Name,
		Email:     user.Email,
		Password:  hashedPassword,
		Role:      user.Role,
		CreatedAt: time.Now(),
	}

	// create to db
	if err := config.DB.Create(&param).Error; err != nil {
		utils.Logger.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    500,
			"message": err.Error(),
		})
	}

	// return success
	utils.Logger.Info("Register successfully")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"message": "Register successfully",
		"data":    nil,
	})
}

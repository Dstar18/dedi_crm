package controllsers

import (
	"dedi_crm/config"
	"dedi_crm/models"
	"dedi_crm/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func Customers(c echo.Context) error {
	var customerM []models.Lead
	config.DB.Find(&customerM)

	// return success
	utils.Logger.Info("Customers successfully")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Customers successfully",
		"data":    customerM,
	})
}

type CustomerValidate struct {
	Name    string `json:"name" validate:"required,min=2,max=50"`
	Email   string `json:"email" validate:"required,email"`
	Phone   string `json:"phone" validate:"required,min=10,max=13"`
	Address string `json:"address" validate:"required,min=2,max=50"`
}

func CustomerStore(c echo.Context) error {
	// request struct validation
	var customer CustomerValidate

	// request params, and check body
	if err := c.Bind(&customer); err != nil {
		utils.Logger.Error("Invalid request body")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "invalid request body",
		})
	}

	// validation struc
	validate := validator.New()
	if err := validate.Struct(customer); err != nil {
		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = "This field is" + " " + err.Tag() + " " + err.Param()
		}
		utils.Logger.Error(errors)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": errors,
		})
	}

	// request struct model
	var customerM models.Lead

	// check email is ready
	if err := config.DB.Where("email = ? ", customer.Email).First(&customerM).Error; err == nil {
		utils.Logger.Warn("email " + customer.Email + " is already")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":   http.StatusBadRequest,
			"mesage": "email " + customer.Email + " is already",
		})
	}

	// chek islogin
	emailIsLogin := c.Get("email").(string)
	var userM models.User
	config.DB.Where("email = ? ", emailIsLogin).First(&userM)

	param := models.Lead{
		Name:      customer.Name,
		Email:     customer.Email,
		Phone:     customer.Phone,
		Address:   customer.Address,
		Status:    "new",
		CreatedBy: userM.ID,
	}

	// create to db
	if err := config.DB.Create(&param).Error; err != nil {
		utils.Logger.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// return success
	utils.Logger.Info("Created successfully")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Created successfully",
	})
}

package controllsers

import (
	"dedi_crm/config"
	"dedi_crm/models"
	"dedi_crm/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type AuthValidate struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func Login(c echo.Context) error {
	// request struct validation
	var user AuthValidate

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

	// check user on database
	var userM models.User
	if err := config.DB.Where("email = ?", user.Email).First(&userM).Error; err != nil {
		utils.Logger.Warn("Invalid email or password")
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"code":    401,
			"message": "Invalid email or password",
		})
	}

	// check password in hash
	if err := utils.CheckPassword(userM.Password, user.Password); err != nil {
		utils.Logger.Warn(err.Error())
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"code":    401,
			"message": err.Error(),
		})
	}

	// Set session
	session, _ := config.Store.Get(c.Request(), "session")
	session.Values["authenticated"] = true
	session.Values["email"] = userM.Email
	session.Save(c.Request(), c.Response())

	// return success
	utils.Logger.Info("Login successfully")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"message": "Login successfully",
		"data":    nil,
	})
}

func Logout(c echo.Context) error {
	session, _ := config.Store.Get(c.Request(), "session")
	session.Options.MaxAge = -1
	session.Save(c.Request(), c.Response())

	// return success
	utils.Logger.Info("Logout successfully")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"message": "Logout successfully",
	})
}

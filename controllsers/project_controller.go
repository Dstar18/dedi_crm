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

type StoreValidate struct {
	LeadID    uint   `json:"lead_id" validate:"required"`
	ProductID string `json:"product_id" validate:"required"`
}

func ProjectAdd(c echo.Context) error {
	// request struct validation
	var project StoreValidate

	// request params, and check body
	if err := c.Bind(&project); err != nil {
		utils.Logger.Error("Invalid request body")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "invalid request body",
		})
	}

	// validation struc
	validate := validator.New()
	if err := validate.Struct(project); err != nil {
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

	// chek islogin
	emailIsLogin := c.Get("email").(string)
	var userM models.User
	config.DB.Where("email = ? ", emailIsLogin).First(&userM)

	param := models.Project{
		LeadID:     project.LeadID,
		ProductID:  project.ProductID,
		Status:     "pending",
		CreatedAt:  time.Now(),
		AssignedTo: userM.ID,
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
	utils.Logger.Info("Project Created successfully")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Project Created successfully",
	})
}

type UpdateValidate struct {
	Status string `json:"status" validate:"required"`
}

func ProjectVerifier(c echo.Context) error {
	// request param id
	id := c.Param("id")

	// request struct model
	var projectM models.Project

	// check data by id
	if err := config.DB.First(&projectM, id).Error; err != nil {
		utils.Logger.Warn("Data Not found")
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    http.StatusNotFound,
			"message": "Data Not Found",
		})
	}

	// request struct validation
	var project UpdateValidate

	// request params, and check body
	if err := c.Bind(&project); err != nil {
		utils.Logger.Error("Invalid request body")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "invalid request body",
		})
	}

	// validation struc
	validate := validator.New()
	if err := validate.Struct(project); err != nil {
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

	// chek islogin
	emailIsLogin := c.Get("email").(string)
	var userM models.User
	config.DB.Where("email = ? ", emailIsLogin).First(&userM)

	// save to db
	projectM.Status = project.Status
	projectM.ApprovalBy = userM.ID
	projectM.ApprovedAt = time.Now()

	if err := config.DB.Save(&projectM).Error; err != nil {
		utils.Logger.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// return success
	utils.Logger.Info("Project Verifier successfully")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Project Verifier successfully",
		"data":    projectM,
	})
}

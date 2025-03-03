package controllsers

import (
	"dedi_crm/config"
	"dedi_crm/models"
	"dedi_crm/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ProductValidate struct {
	Name        string `json:"name" validate:"required,min=1,max=50"`
	Desciprtion string `json:"description" validate:"required,max=50"`
	Price       string `json:"price" validate:"required"`
}

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

func ProductStore(c echo.Context) error {
	// request struct validation
	var product ProductValidate

	// request params, and check body
	if err := c.Bind(&product); err != nil {
		utils.Logger.Error("Invalid request body")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "invalid request body",
		})
	}

	// validation struc
	validate := validator.New()
	if err := validate.Struct(product); err != nil {
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
	var productM models.Product

	// check name is ready
	if err := config.DB.Where("name = ? ", product.Name).First(&productM).Error; err == nil {
		utils.Logger.Warn("name " + product.Name + " is already")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":   http.StatusBadRequest,
			"mesage": "name " + product.Name + " is already",
		})
	}

	param := models.Product{
		Name:        product.Name,
		Desciprtion: product.Desciprtion,
		Price:       product.Price,
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

func ProductUpdate(c echo.Context) error {
	// request param id
	id := c.Param("id")

	// request struct model
	var productM models.Product

	// check data by id
	if err := config.DB.First(&productM, id).Error; err != nil {
		utils.Logger.Warn("Data Not found")
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    http.StatusNotFound,
			"message": "Data Not Found",
		})
	}

	// request struct validation
	var product ProductValidate

	// request params, and check body
	if err := c.Bind(&product); err != nil {
		utils.Logger.Error("Invalid request body")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "invalid request body",
		})
	}

	// validation struc
	validate := validator.New()
	if err := validate.Struct(product); err != nil {
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

	// save to db
	productM.Name = product.Name
	productM.Desciprtion = product.Desciprtion
	productM.Price = product.Price

	if err := config.DB.Save(&productM).Error; err != nil {
		utils.Logger.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// return success
	utils.Logger.Info("Update successfully")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Update successfully",
		"data":    productM,
	})
}

func ProductDestroy(c echo.Context) error {
	// request param id
	id := c.Param("id")

	// request struct model
	var productM models.Product

	// check data by id
	if err := config.DB.First(&productM, id).Error; err != nil {
		utils.Logger.Warn("Data Not found")
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    http.StatusNotFound,
			"message": "Data Not Found",
		})
	}

	// delete to db
	if err := config.DB.Delete(&productM).Error; err != nil {
		utils.Logger.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// return success
	utils.Logger.Info("Deleted successfully")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Deleted successfully",
	})
}

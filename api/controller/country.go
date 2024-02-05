package controller

import (
	"database/sql"

	"net/http"

	"github.com/gin-gonic/gin"
)

type CountryController struct {
	Usecase country.Usecase
	Env     *config.Env
}

func (cc *CountryController) ListWithPrice(c *gin.Context) {
	countries, err := cc.Usecase.List(c)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, countries)
}

// func (cc *CountryController) Create(c *gin.Context) {
// 	var req dto.InsertCountryReq

// 	err := c.ShouldBind(&req)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
// 	}

// 	arg := db.CreateCountryParams{}

// 	country, err := cc.Usecase.Insert()
// }

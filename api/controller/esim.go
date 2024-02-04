package controller

import (
	"database/sql"
	"esim-service/config"
	"esim-service/domain"
	"esim-service/domain/country"
	"esim-service/domain/esim"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EsimController struct {
	Usecase esim.Usecase
	Env     *config.Env
}

func (ec *EsimController) ListByCountry(c *gin.Context) {
	var req country.ListByCountryReq

	err := c.ShouldBindQuery(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	countries, err := ec.Usecase.ListByCountry(c, req.CountryCode)
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

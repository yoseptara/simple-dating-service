package controller

import (
	"database/sql"
<<<<<<< HEAD

=======
	"esim-service/config"
	"esim-service/domain"
>>>>>>> 5cc17bee3d7dc905b0f996385046d65a42c4841d
	"net/http"
	country "simple-dating-app-service/domain/user"

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

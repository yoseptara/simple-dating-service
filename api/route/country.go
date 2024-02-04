package route

import (
	"esim-service/api/controller"
	"esim-service/domain"
	"esim-service/usecase"

	"github.com/gin-gonic/gin"
)

func NewCountryRouter(server *domain.ConcreteServer, group *gin.RouterGroup) {
	cc := &controller.CountryController{
		Usecase: usecase.NewCountryUsecase(server),
	}

	group.GET("/", cc.ListWithPrice)
}

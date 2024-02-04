package route

import (
	"esim-service/api/controller"
	"esim-service/domain"
	"esim-service/usecase"

	"github.com/gin-gonic/gin"
)

func NewEsimRouter(server *domain.ConcreteServer, group *gin.RouterGroup) {
	ec := &controller.EsimController{
		Usecase: usecase.NewEsimUsecase(server),
	}

	group.GET("/", ec.ListByCountry)
}

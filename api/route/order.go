package route

import (
	"esim-service/api/controller"
	"esim-service/domain"
	repository "esim-service/repository/http"
	"esim-service/service"
	"esim-service/usecase"

	"github.com/gin-gonic/gin"
)

func NewOrderRouter(server *domain.ConcreteServer, group *gin.RouterGroup) {
	xhr := repository.NewXenditHttpRepository(server.Config)
	uhr := repository.NewUsimsaHttpRepository(server.Config)
	ss := service.NewSmtpService(server.Config)
	oc := &controller.OrderController{
		Usecase: usecase.NewOrderUsecase(server, ss, xhr, uhr),
	}

	group.POST("/", oc.CreateOrder)
	group.POST("/xendit_invoice_webhook", oc.HandleOrderInvoicePaymentCallback)
	group.POST("/usimsa_subscribe_order_webhook", oc.HandleUsimsaSubscribeOrderCallback)
}

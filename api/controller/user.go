package controller

import (
	"database/sql"
	"esim-service/config"
	"esim-service/domain"
	"esim-service/domain/order"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Usecase order.Usecase
	Env     *config.Env
}

func (oc *UserController) CreateOrder(c *gin.Context) {
	var req order.CreateOrderReq

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	order, err := oc.Usecase.CreatePendingOrder(c, req.EsimId, req.Quantity, req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}

func (oc *UserController) HandleOrderInvoicePaymentCallback(c *gin.Context) {
	var req order.InvoicePaymentCallback

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	order, err := oc.Usecase.UpdateInvoicePaymentAndOrderUsimsa(c, req)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}

func (oc *UserController) HandleUsimsaSubscribeOrderCallback(c *gin.Context) {
	if c.Request.Method == http.MethodPost {
		var req order.UsimsaSubscribedOrderReq

		err := c.ShouldBindJSON(&req)
		if err != nil {
			c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
			return
		}

		order, err := oc.Usecase.SendPurchasedEsimEmail(c, req.TopupId, req.Iccid, req.Smdp, req.ActivateCode, req.QrcodeImgUrl)
		if err != nil {
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
			return
		}

		c.JSON(http.StatusOK, order)

	} else {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"message": "Invalid request method",
		})
	}
}

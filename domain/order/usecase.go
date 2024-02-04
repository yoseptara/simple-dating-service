package order

import (
	"context"
	db "esim-service/db/sqlc"

	"github.com/aws/aws-sdk-go/service/sesv2"
)

type Usecase interface {
	CreatePendingOrder(c context.Context, esimId int64, quantity int32, email string) (*CreateOrderResp, error)
	UpdateInvoicePaymentAndOrderUsimsa(c context.Context, callback InvoicePaymentCallback) (*db.Order, error)
	SendPurchasedEsimEmail(c context.Context, topupId string, iccid string, smdp string, activationCode string, qrCodeImgUrl string) (*sesv2.SendEmailOutput, error)
	Get(c context.Context, id int64) (*db.Order, error)
}

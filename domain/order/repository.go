package order

type XenditHttpRepository interface {
	CreateInvoice(req CreateOrderInvoiceReq) (*CreateOrderInvoiceResp, error)
}

type UsimsaHttpRepository interface {
	SubscribeOrder(req SubscribeUsimsaOrderReq) (*SubscribeUsimsaOrderResp, error)
}

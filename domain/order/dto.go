package order

import "esim-service/domain"

type CreateOrderReq struct {
	EsimId   int64  `form:"esim_id" json:"esim_id" binding:"required"`
	Quantity int32  `form:"quantity" json:"quantity" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
}

type CreateOrderResp struct {
	InvoiceUrl string `json:"invoice_url" binding:"required"`
}

type SubscribeUsimsaOrderReq struct {
	UniqueOrderId string                   `form:"order_id" json:"orderId" binding:"required"`
	Products      []domain.UsimsaOrderItem `form:"products" json:"products" binding:"required"`
}

type SubscribeUsimsaOrderResp struct {
	Code     string                     `form:"code" json:"code" binding:"required"`
	Message  string                     `form:"message" json:"message" binding:"required"`
	Products []domain.UsimsaOrderedItem `form:"products" json:"products" binding:"required"`
}

type UsimsaSubscribedOrderReq struct {
	TopupId      string `json:"topupId" binding:"required"`
	OptionId     string `json:"optionId" binding:"required"`
	Iccid        string `json:"iccid" binding:"required"`
	Smdp         string `json:"smdp" binding:"required"`
	ActivateCode string `json:"activateCode" binding:"required"`
	DownloadLink string `json:"downloadLink" binding:"required"`
	QrcodeImgUrl string `json:"qrcodeImgUrl" binding:"required"`
}

type CreateOrderInvoiceReq struct {
	ExternalID                     string                               `json:"external_id" binding:"required"`
	Amount                         int                                  `json:"amount" binding:"required"`
	Description                    string                               `json:"description,omitempty"`
	InvoiceDuration                int                                  `json:"invoice_duration,omitempty"`
	Customer                       domain.InvoiceCustomer               `json:"customer,omitempty"`
	CustomerNotificationPreference domain.InvoiceNotificationPreference `json:"customer_notification_preference,omitempty"`
	SuccessRedirectURL             string                               `json:"success_redirect_url,omitempty"`
	FailureRedirectURL             string                               `json:"failure_redirect_url,omitempty"`
	Currency                       string                               `json:"currency,omitempty"`
	Items                          []domain.InvoiceItem                 `json:"items,omitempty"`
	Fees                           []domain.InvoiceFee                  `json:"fees,omitempty"`
}

type CreateOrderInvoiceResp struct {
	ID                        string                   `json:"id"`
	UserID                    string                   `json:"user_id"`
	ExternalID                string                   `json:"external_id"`
	Status                    string                   `json:"status"`
	MerchantName              string                   `json:"merchant_name"`
	MerchantProfilePictureURL string                   `json:"merchant_profile_picture_url"`
	Amount                    int                      `json:"amount"`
	PayerEmail                string                   `json:"payer_email"`
	Description               string                   `json:"description"`
	InvoiceURL                string                   `json:"invoice_url"`
	ExpiryDate                string                   `json:"expiry_date"`
	AvailableBanks            []domain.InvoiceBank     `json:"available_banks"`
	AvailableRetailOutlets    []domain.InvoiceRetail   `json:"available_retail_outlets"`
	AvailablePaylaters        []domain.InvoicePaylater `json:"available_paylaters"`
	AvailableQRCodes          []domain.InvoiceQRCode   `json:"available_qr_codes"`
	ShouldExcludeCreditCard   bool                     `json:"should_exclude_credit_card"`
	ShouldSendEmail           bool                     `json:"should_send_email"`
	Created                   string                   `json:"created"`
	Updated                   string                   `json:"updated"`
	MIDLabel                  string                   `json:"mid_label"`
	Currency                  string                   `json:"currency"`
	FixedVA                   bool                     `json:"fixed_va"`
	Locale                    string                   `json:"locale"`
	Customer                  domain.InvoiceCustomer   `json:"customer"`
	Items                     []domain.InvoiceItem     `json:"items"`
	Fees                      []domain.InvoiceFee      `json:"fees"`
}

type InvoicePaymentCallback struct {
	ID                     string                `json:"id"`
	ExternalID             string                `json:"external_id"`
	UserID                 string                `json:"user_id"`
	IsHigh                 bool                  `json:"is_high"`
	Status                 string                `json:"status"`
	MerchantName           string                `json:"merchant_name"`
	Amount                 float64               `json:"amount"`
	PayerEmail             string                `json:"payer_email,omitempty"`
	Description            string                `json:"description,omitempty"`
	PaidAmount             float64               `json:"paid_amount,omitempty"`
	Updated                string                `json:"updated"`
	Created                string                `json:"created"`
	Currency               string                `json:"currency,omitempty"`
	PaidAt                 string                `json:"paid_at,omitempty"`
	PaymentMethod          string                `json:"payment_method,omitempty"`
	PaymentChannel         string                `json:"payment_channel,omitempty"`
	PaymentDestination     string                `json:"payment_destination,omitempty"`
	PaymentDetails         domain.InvoiceDetails `json:"payment_details,omitempty"`
	PaymentID              string                `json:"payment_id,omitempty"`
	SuccessRedirectURL     string                `json:"success_redirect_url,omitempty"`
	FailureRedirectURL     string                `json:"failure_redirect_url,omitempty"`
	CreditCardChargeID     string                `json:"credit_card_charge_id,omitempty"`
	Items                  []domain.InvoiceItem  `json:"items,omitempty"`
	Fees                   []domain.InvoiceFee   `json:"fees,omitempty"`
	ShouldAuthenticateCard bool                  `json:"should_authenticate_credit_card,omitempty"`
	BankCode               string                `json:"bank_code,omitempty"`
	EWalletType            string                `json:"ewallet_type,omitempty"`
	OnDemandLink           string                `json:"on_demand_link,omitempty"`
	RecurringPaymentID     string                `json:"recurring_payment_id,omitempty"`
	FeesPaidAmount         float64               `json:"fees_paid_amount,omitempty"`
	AdjustedReceivedAmount float64               `json:"adjusted_received_amount,omitempty"`
}

package domain

type InvoiceDetails struct {
	ReceiptID string `json:"receipt_id,omitempty"`
	Source    string `json:"source,omitempty"`
}

type InvoiceCustomer struct {
	GivenNames   string           `json:"given_names,omitempty"`
	Surname      string           `json:"surname,omitempty"`
	Email        string           `json:"email,omitempty"`
	MobileNumber string           `json:"mobile_number,omitempty"`
	Addresses    []InvoiceAddress `json:"addresses,omitempty"`
}

type InvoiceAddress struct {
	City        string `json:"city,omitempty"`
	Country     string `json:"country,omitempty"`
	PostalCode  string `json:"postal_code,omitempty"`
	State       string `json:"state,omitempty"`
	StreetLine1 string `json:"street_line1,omitempty"`
	StreetLine2 string `json:"street_line2,omitempty"`
}

type InvoiceNotificationPreference struct {
	InvoiceCreated  []string `json:"invoice_created,omitempty"`
	InvoiceReminder []string `json:"invoice_reminder,omitempty"`
	InvoicePaid     []string `json:"invoice_paid,omitempty"`
	InvoiceExpired  []string `json:"invoice_expired,omitempty"`
}

type InvoiceItem struct {
	Name     string `json:"name" binding:"required"`
	Quantity int    `json:"quantity" binding:"required"`
	Price    int    `json:"price" binding:"required"`
	Category string `json:"category,omitempty"`
	URL      string `json:"url,omitempty"`
}

type InvoiceFee struct {
	Type  string `json:"type" binding:"required"`
	Value int    `json:"value" binding:"required"`
}

type InvoiceBank struct {
	BankCode          string `json:"bank_code,omitempty"`
	CollectionType    string `json:"collection_type,omitempty"`
	TransferAmount    int    `json:"transfer_amount,omitempty"`
	BankBranch        string `json:"bank_branch,omitempty"`
	AccountHolderName string `json:"account_holder_name,omitempty"`
}

type InvoiceRetail struct {
	RetailOutletName string `json:"retail_outlet_name,omitempty"`
}

type InvoicePaylater struct {
	PaylaterType string `json:"paylater_type,omitempty"`
}

type InvoiceQRCode struct {
	QRCodeType string `json:"qr_code_type,omitempty"`
}

package domain

type UsimsaOrderItem struct {
	OptionId string `form:"optionId" json:"optionId" binding:"required"`
	Quantity int32  `form:"qty" json:"qty" binding:"required"`
}

type UsimsaOrderedItem struct {
	TopupId  string `form:"topupId" json:"topupId" binding:"required"`
	OptionId string `form:"optionId" json:"optionId" binding:"required"`
}

package esim

type CreateEsimReq struct {
	CountryCode string `json:"country_code" binding:"required"`
}

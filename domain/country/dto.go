package country

type InsertCountryReq struct {
	CountryCode string `json:"country_code" binding:"required,iso3166_1_alpha2"`
	Name        string `json:"name" binding:"required"`
}

type ListByCountryReq struct {
	CountryCode string `form:"country_code" json:"country_code" binding:"required,country_code"`
}

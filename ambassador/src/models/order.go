package models

type Order struct {
	Id     uint    `json:"id" gorm:"primary_key; autoIncrement:false"`
	UserId uint    `json:"user_id"`
	Code   string  `json:"code"`
	Total  float64 `json:"total"`
}

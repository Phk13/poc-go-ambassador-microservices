package models

type Product struct {
	Id          uint    `json:"id" gorm:"primary_key; autoIncrement:false"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"`
}

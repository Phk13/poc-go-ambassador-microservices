package models

import (
	"gorm.io/gorm"
)

type User struct {
	Id           uint    `json:"id"`
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	Email        string  `json:"email"`
	Password     []byte  `json:"-"`
	IsAmbassador bool    `json:"is_ambassador"`
	Revenue      float64 `json:"revenue,omitempty"`
}

func (user *User) Name() string {
	return user.FirstName + " " + user.LastName
}

type Ambassador User

func (ambassador *Ambassador) CalculateRevenue(db *gorm.DB) {
	var orders []Order

	db.Find(&orders, &Order{
		UserId: ambassador.Id,
	})

	var revenue float64

	for _, order := range orders {
		revenue += order.Total
	}

	ambassador.Revenue = revenue
}

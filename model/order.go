package model

// import (
// 	"github.com/jinzhu/gorm"
// )
	

type Order struct {
	// gorm.Model
	Order_id string	`gorm:"size:50"`
	Merchant_id string `gorm:"size:50"`
	User_id int `gorm:"size:11"`
	Order_date string `gorm:"size:11"`
	Order_total int `gorm:"size:11"`
	Order_return int `gorm:"size:11"`
	Order_status int `gorm:"size:11"`
	Created_at string `gorm:"size:15"`
	Updated_at string `gorm:"size:15"`
	Deleted_at string `gorm:"size:15"`
}
package model

import (
	"github.com/jinzhu/gorm"
)

type Order struct {
	// gorm.Model
	Id           int `gorm:"AUTO_INCREMENT"`
	Merchant_id  int
	User_id      int    `gorm:"size:11"`
	Order_date   string `gorm:"size:11"`
	Order_total  int    `gorm:"size:11"`
	Order_return int    `gorm:"size:11"`
	Order_status int    `gorm:"size:11"`
	Created_at   string `gorm:"size:15"`
	Updated_at   string `gorm:"size:15"`
	// Deleted_at   string `gorm:"size:15"`
	Deleted_at string `gorm:"DEFAULT:NULL"` // ignore this field

}

func (o *Order) AddOrder(db *gorm.DB) []error {
	errors := db.Create(&o).GetErrors()
	return errors
}

func (o *Order) GetOrder(db *gorm.DB) []error {
	errors := db.Find(&o).GetErrors()
	return errors
}

func (o *Order) UpdateOrder(db *gorm.DB) []error {
	errors := db.First(&o).GetErrors()
	if len(errors) > 0 {
		return errors
	}
	errors = db.Save(&o).GetErrors()
	return errors
}

func (o *Order) DeleteOrder(db *gorm.DB) []error {
	errors := db.Delete(&o).GetErrors()
	return errors
}

func GetAllOrder(db *gorm.DB) ([]Order, []error) {
	var orders []Order
	errors := db.Find(&orders).GetErrors()
	return orders, errors
}

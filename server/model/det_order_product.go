package model

import (
	// "fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type Det_order_product struct {
	Id                         int        `gorm:"AUTO_INCREMENT"`
	Order_id                   string     `gorm:"size:50"`
	Product_id                 int        `gorm:"size:11"`
	Det_order_product_qty      int        `gorm:"size:11"`
	Det_order_product_total    int        `gorm:"size:50"`
	Det_order_product_discount int        `gorm:"size:50"`
	Det_order_product_price    int        `gorm:"size:50"`
	Det_order_product_status   int        `gorm:"size:11"`
	Created_at                 *time.Time `gorm:"type:timestamp"`
	Updated_at                 *time.Time
	Deleted_at                 *time.Time
}

func (dop *Det_order_product) AddDetOrder(db *gorm.DB) []error {
	db.Create(&dop)
	errors := db.Create(dop).GetErrors()
	return errors
}

func (dop *Det_order_product) GetDetOrder(db *gorm.DB) []error {
	errors := db.Find(&dop).GetErrors()
	return errors
}

func (dop *Det_order_product) UpdateDetOrder(db *gorm.DB) []error{
	errors := db.First(&dop).GetErrors()
	if len(errors)>0 {
		return errors
	}
	errors = db.Save(&dop).GetErrors()
	return errors
}

func (dop *Det_order_product) DeleteDetOrder(db *gorm.DB) []error{
	errors := db.First(&dop).GetErrors()
	if len(errors)>0 {
		return errors
	}
	res := db.Delete(&dop)
	return res.GetErrors()
}

func GetAllDetOrder (db *gorm.DB) ([]Det_order_product, []error){
	var det_order_products []Det_order_product
	res := db.Find(det_order_products)
	return det_order_products, res.GetErrors()
}
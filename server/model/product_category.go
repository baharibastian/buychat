package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Product_category struct {
	Id int `gorm:"AUTO_INCREMENT"`
	Product_category_name string `gorm:"size:50"`
	Product_category_description string `gorm:"size:255"`
	Created_at *time.Time `gorm:"type:timestamp"`
	Updated_at *time.Time 
	Deleted_at   *time.Time
}

func (p *Product_category) AddCategory(db *gorm.DB) bool {
	db.Create(&p)
	err := db.NewRecord(p)
	return err
}

func (p *Product_category) DeleteCategory(db *gorm.DB) (map[string]interface{}, []error) {
	err := db.First(&p)
	message_success := map[string]interface{} {
		"Status" : "Delete Success",
	}
	f := map[string]interface{} {
		"Name" : p.Product_category_name,
		"Description" : p.Product_category_description,
		"Message" : message_success,
	}
	if len(err.GetErrors()) > 0 {
		return nil, err.GetErrors()
	}
	err = db.Delete(&p)
	if len(err.GetErrors()) > 0 {
		return nil, err.GetErrors()
	}
	return f, nil
}

func (p *Product_category) UpdateProduct(db *gorm.DB) []error {
	db.First(&p)
	err := db.Save(&p)
	return err.GetErrors()
}

func GetAllProductCategory(db *gorm.DB) ([]Product_category, []error) {
	var product_categories []Product_category
	res := db.Find(&product_categories)
	return product_categories, res.GetErrors()
}
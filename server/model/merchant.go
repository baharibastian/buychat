package model

import (
	"github.com/jinzhu/gorm"
	"fmt"
)

type Merchant struct {
	Id                   int    `gorm:"AUTO_INCREMENT"`
	Merchant_code        string `gorm:"size:50"`
	Merchant_name        string `gorm:"size:100"`
	Merchant_description string `gorm:"size:255"`
	Merchant_status      int    `gorm:"size:11"`
	Merchant_open        string `gorm:"size:11"`
	Merchant_close       string `gorm:"size:11"`
	Created_at           string `gorm:"size:15"`
	Updated_at           string `gorm:"size:15"`
	Deleted_at           string `gorm:"size:15"`
	Products 			 []Product `gorm:"foreignkey:Merchant_Id`

}

func (p *Merchant) AddMerchant(db *gorm.DB) bool {
	db.Create(&p)
	err := db.NewRecord(p) //if success, will return false otherwise true
	return err
}
 func (p *Merchant) DeleteMerchant(db *gorm.DB) (map[string]interface{}, []error) {
	err := db.First(&p)
	message_success := map[string]interface{}{
		"Status": "Delete Success",
	}
	f := map[string]interface{}{
		"Merchant Code": p.Merchant_code,
		"Merchant Name": p.Merchant_name,
		"Message":       message_success,
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
 func (p *Merchant) UpdateMerchant(db *gorm.DB) []error {
	db.First(&p)
	err := db.Save(&p)
	return err.GetErrors()
}
 func GetAllMerchants(db *gorm.DB) ([]Merchant, []error) {
	var merchants []Merchant
	errors := db.Find(&merchants).GetErrors()
	if len(errors) > 0 {
		return nil, errors
	}
	var new_merchants []Merchant
	for _, merch := range merchants {
		fmt.Println(merch)
		var merchant Merchant
		fmt.Println(merch.Id)
		db.Find(&merch).Related(&merch.Products) //SELECT * FROM merchants JOIN products ON merchants.id = products.merchant_id WHERE merchant_id = merch.id
		merchant = merch
		new_merchants = append(new_merchants, merchant)
	}
	return new_merchants, errors
}
 func (p *Merchant) GetMerchant(db *gorm.DB) []error {
	err := db.First(&p)
	return err.GetErrors()
 }
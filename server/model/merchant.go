package model

import (
<<<<<<< HEAD
	//"fmt"
=======
	"fmt"
>>>>>>> 70a94b4de082a58cc8f742cb6dd36f725068e634
	"time"

	"github.com/jinzhu/gorm"
)

type Merchant struct {
	Id                   int        `gorm:"AUTO_INCREMENT"`
	Merchant_code        string     `gorm:"size:50"`
	Merchant_name        string     `gorm:"size:100"`
	Merchant_description string     `gorm:"size:255"`
	Merchant_status      int        `gorm:"size:11"`
	Merchant_open        string     `gorm:"size:11"`
	Merchant_close       string     `gorm:"size:11"`
	Created_at           *time.Time `gorm:"type:timestamp"`
	Updated_at           *time.Time
	Deleted_at           *time.Time
	Products             []Product `gorm:"foreignkey:Merchant_Id`
}

func (p *Merchant) AddMerchant(db *gorm.DB) []error {
	errors := db.Create(&p).GetErrors()
	return errors
}

func (p *Merchant) DeleteMerchant(db *gorm.DB) (map[string]interface{}, []error) {
	errors := db.First(&p).GetErrors()

	if len(errors) > 0 {
		return nil, errors
	}
	errors = db.Delete(&p).GetErrors()

	message_success := map[string]interface{}{
		"Status": "Delete Success",
	}

	f := map[string]interface{}{
		"Merchant Code": p.Merchant_code,
		"Merchant Name": p.Merchant_name,
		"Message":       message_success,
	}
	return f, errors
}
func (p *Merchant) UpdateMerchant(db *gorm.DB) []error {
	errors := db.First(&p).GetErrors()
	if len(errors) > 0 {
		return errors
	}
	errors = db.Save(&p).GetErrors()
	return errors
}

func GetAllMerchants(db *gorm.DB) ([]Merchant, []error) {
	var merchants []Merchant
	errors := db.Find(&merchants).GetErrors()
	if len(errors) > 0 {
		return nil, errors
	}
	var new_merchants []Merchant
	for _, merch := range merchants {
		var merchant Merchant
		db.Find(&merch).Related(&merch.Products) //SELECT * FROM merchants JOIN products ON merchants.id = products.merchant_id WHERE merchant_id = merch.id
		merchant = merch
		fmt.Println("nohh")
		fmt.Println(&merch)
		new_merchants = append(new_merchants, merchant)
	}

	return new_merchants, errors
}

func (p *Merchant) GetMerchant(db *gorm.DB) []error {
	errors := db.First(&p).GetErrors()
	return errors
}

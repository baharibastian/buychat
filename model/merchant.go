package model 

type Merchant struct {
	Merchant_id int `gorm:"AUTO_INCREMENT"`
	Merchant_code string `gorm:"size:50"`
	Merchant_name string `gorm:"size:100"`
	Merchant_description string `gorm:"size:255"`
	Merchant_status int `gorm:"size:11"`
	Merchant_open string `gorm:"size:11"`
	Merchant_close string `gorm:"size:11"`
	Created_at string `gorm:"size:15"`
	Updated_at string `gorm:"size:15"`
	Deleted_at string `gorm:"size:15"`
}
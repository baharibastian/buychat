package model

type Det_order_product struct {
	Id                         int    `gorm:"AUTO_INCREMENT"`
	Order_id                   string `gorm:"size:50"`
	Product_id                 int    `gorm:"size:11"`
	Det_order_product_qty      int    `gorm:"size:11"`
	Det_order_product_total    int    `gorm:"size:50"`
	Det_order_product_discount int    `gorm:"size:50"`
	Det_order_product_price    int    `gorm:"size:50"`
	Det_order_product_status   int    `gorm:"size:11"`
	Created_at                 string `gorm:"size:15"`
	Updated_at                 string `gorm:"size:15"`
}

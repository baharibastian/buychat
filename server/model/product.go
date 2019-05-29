package model

import (
	"time"
	// "database/sql"

	"github.com/jinzhu/gorm"

	// "fmt"
)

type Product struct {
	Id int `gorm:"AUTO_INCREMENT"`
	Merchant_id int `gorm:"size:15"`
	Product_code string `gorm:"size:50"`
	Product_name string `gorm:"size:100"`
	Product_category_id int `gorm:"size:11"`
	Product_description string `gorm:"size:255"`
	Product_base_price int `gorm:"size:15"`
	Product_sell_price int `gorm:"size:15"`
	Created_at *time.Time `gorm:"type:timestamp"`
	Updated_at *time.Time 
	Deleted_at   *time.Time
	Merchant Merchant `gorm:"foreignkey:Merchant_id`
}

// var err error

func (p *Product) AddProduct(db *gorm.DB) bool {
	// statement := fmt.Sprintf("INSERT INTO PRODUCTS (NAME, PRICE, USER_ID) Values('%s',%d,%d)", p.Name, p.Price, p.UserId)
	// res, err := db.Exec(statement)
	// if err != nil {
	// 	return err
	// }
	// id, err := res.LastInsertId()
	// p.Id = int(id)

	// db.NewRecord(p) // => returns `true` as primary key is blank
	db.Create(&p)
	err := db.NewRecord(p) // => return `false` after `user` created
	return err

}

func (p *Product) DeleteProduct(db *gorm.DB) (map[string]interface{}, []error) {
	// statement := fmt.Sprintf("DELETE PRODUCTS WHERE ID= %d", p.Id)
	// _, err := db.Exec(statement)
	// return err
	// err := db.Delete(p).GetErrors()
	// var new_err error
	// err = db.Where("id=%d",p.Id).Find(&p).GetErrors()
	// statement := fmt.Sprintf("id=%d",p.Id)
	// var test map[string]interface{}
	// var test = make(map[string]interface{})
	// err := db.Where(statement).Scan(test)
	
	// err := db.First(&p, p.Id)
	
	err := db.First(&p)
	message_success := map[string]interface{}{
		"Status" : "Delete Success",
	}
	f:= map[string]interface{}{
		"Name" : p.Product_name,
		"Price" : p.Product_sell_price,
		"Message" : message_success,
	}
	if len(err.GetErrors())>0 {
		return nil, err.GetErrors()
	}
	err = db.Delete(&p)
	if len(err.GetErrors())>0 {
		return nil, err.GetErrors()
	}
	return f, nil
}

// func (p *Product) UpdateProduct(db *sql.DB) error {
func (p *Product) UpdateProduct(db *gorm.DB) []error {
	// statement := fmt.Sprintf("UPDATE PRODUCTS SET NAME='%s' Price= %d where id=%d", p.Name, p.Price, p.Id)
	// _, err := db.Exec(statement)
	// return err
	db.First(&p)
	err := db.Save(&p)
	return err.GetErrors()	
}

func GetAllProduct(db *gorm.DB, id int) ([]Product, []error) {
	var products []Product
	var new_products []Product
	if id > 0 {
		errors := db.Where("product_category_id = ?", id).Find(&products).GetErrors()
		if len(errors) > 0 {
			return nil, errors
		}
		for _, product := range products {
			var prod Product
			db.Find(&product).Related(&product.Merchant)
			prod = product
			new_products = append(new_products, prod)
		}
		res := db
		return new_products, res.GetErrors()
	} else {
		res := db.Find(&products)
		return products, res.GetErrors()
	}
	
}

func (p *Product) GetProduct(db *gorm.DB) []error {
	// statement := fmt.Sprintf("SELECT * FROM PRODUCTS WHERE ID = %d", p.Id)
	// err := db.QueryRow(statement).Scan(&p.Id, &p.Name, &p.Price, &p.User_Id)
	// return err
	err := db.First(&p)
	return err.GetErrors()
}

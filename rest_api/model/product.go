package model

import (
	// "database/sql"

	"github.com/jinzhu/gorm"

	// "fmt"
)

type Product struct {
	// gorm.Model
	Id      int    `gorm:"AUTO_INCREMENT"`
	Name    string `gorm:"size:50"`
	Price   int
	User_Id int
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
		"Name" : p.Id,
		"Price" : p.Name,
		"User_Id" : p.User_Id,
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

func GetAllProduct(db *gorm.DB) ([]Product, []error) {
	// statement := fmt.Sprintf("SELECT * FROM PRODUCTS")
	// var products []Product
	// rows, err := db.Query(statement)
	// for rows.Next() {
	// 	p := Product{}
	// 	err := rows.Scan(&p.Id, &p.Name, &p.Price, &p.User_Id)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	products = append(products, p)
	// }
	// return products, err
	var products []Product
	res := db.Find(&products)
	return products, res.GetErrors()
}

func (p *Product) GetProduct(db *gorm.DB) []error {
	// statement := fmt.Sprintf("SELECT * FROM PRODUCTS WHERE ID = %d", p.Id)
	// err := db.QueryRow(statement).Scan(&p.Id, &p.Name, &p.Price, &p.User_Id)
	// return err
	err := db.First(&p)
	return err.GetErrors()
}

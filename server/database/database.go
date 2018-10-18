package database

import (
	// "fmt"
	"log"

	// "database/sql"
	// _ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/buychat/server/model"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open("mysql", "sql12261256:vthTmGeATF@tcp(sql12.freemysqlhosting.net)/sql12261256?parseTime=true")

	// DB.AutoMigrate(&model.Product{})
	// DB.DropTable(&model.Product{})
	DB.AutoMigrate(&model.Order{})
	DB.AutoMigrate(&model.Merchant{})
	DB.AutoMigrate(&model.Product{})
	DB.Model(&model.Product{}).AddForeignKey("merchant_id", "merchants(id)", "RESTRICT", "RESTRICT")
	DB.AutoMigrate(&model.Det_order_product{})
	DB.Model(&model.Det_order_product{}).AddForeignKey("order_id", "orders(id)", "RESTRICT", "RESTRICT")
	DB.Model(&model.Det_order_product{}).AddForeignKey("product_id", "products(id)", "RESTRICT", "RESTRICT")
	DB.AutoMigrate(&model.Product_category{})
	// DB.AutoMigrate(&model.User{})
	// DB, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang_new")
	if err != nil {
		log.Fatal(err)
	}
}

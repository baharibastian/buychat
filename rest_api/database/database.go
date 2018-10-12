package database

import (
	// "fmt"
	"log"

	// "database/sql"
	// _ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/jinzhu/gorm"
	"rest_api/model"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open("mysql", "root:@/golang_new2")
	// DB.AutoMigrate(&model.Product{})
	// DB.DropTable(&model.Product{})
	DB.AutoMigrate(&model.Product{})
	DB.AutoMigrate(&model.User{})
	// DB, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang_new")
	if err != nil {
		log.Fatal(err)
	}
}

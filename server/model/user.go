package model

import (
	// "database/sql"
	"github.com/jinzhu/gorm"
	// "fmt"
)

type User struct {
	Id       int `gorm:"AUTO_INCREMENT"`
	Name     string `gorm:"size:50"`
	Age      int
	Products []Product `gorm:"foreignkey:User_Id`
}

func (u *User) GetUser(db *gorm.DB) (interface{},[]error) {
	// statement := fmt.Sprintf("SELECT * FROM USERS WHERE ID=%d", u.Id)
	// err:= db.QueryRow(statement).Scan(&u.Id, &u.Name, &u.Age)
	// if err!= nil {
	// 	return nil, err
	// }
	res := db.First(&u)
	if len(res.GetErrors())>0 {
		return nil, res.GetErrors()
	}

	f := map[string]interface{}{
		"Id": u.Id,
		"Name":  u.Name,
		"Age": u.Age,
	}
	
	return f, nil

}

func GetUsers(db *gorm.DB) ([]User, []error) {
	// statement := fmt.Sprintf("SELECT * FROM USERS")
	// rows, err := db.Query(statement)
	// if err != nil {
	// 	return nil, err
	// }
	// defer rows.Close()
	// users := []User{}
	// for rows.Next() {
	// 	var u User
	// 	err = rows.Scan(&u.Id, &u.Name, &u.Age)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	users = append(users, u)
	// }
	// return users, nil
	var users []User
	res := db.Find(&users)
	return users, res.GetErrors()
}

func (u *User) UpdateUser(db *gorm.DB) []error {
	// statement := fmt.Sprintf("UPDATE USERS  SET Name='%s', Age=%d WHERE ID=%d", u.Name, u.Age, u.Id)
	// fmt.Println("db update user")
	// // statement := fmt.Sprintf("UPDATE users SET name='%s', age=%d WHERE id=%d", u.Name, u.Age, u.ID)

	// _, err := db.Exec(statement)
	// return err
	db.First(&u)
	res := db.Save(&u)
	return res.GetErrors()

}

func (u *User) DeleteUser(db *gorm.DB) []error {
	// statement := fmt.Sprintf("DELETE USERS WHERE ID=%d", u.Id)
	// _, err := db.Exec(statement)
	// return err
	errors := db.First(&u).GetErrors()
	if len(errors)>0 {
		return errors
	}
	errors = db.Delete(&u).GetErrors()
	return errors
}

func (u *User) AddUser(db *gorm.DB) []error {
	// statement := fmt.Sprintf("INSERT INTO USERS (name,age) values('%s',%d)", u.Name, u.Age)
	// res, err := db.Exec(statement)
	// if err != nil {
	// 	return err
	// } else {
	// 	id, err := res.LastInsertId()
	// 	if err != nil {
	// 		return err
	// 	}
	// 	u.Id = int(id)
	// 	return err
	// }
	db.NewRecord(u)
	errors := db.Create(&u).GetErrors()
	return errors
}

func (u *User) GetUserProduct(db *gorm.DB)  []error {
	// statement := fmt.Sprintf("SELECT * FROM USERS.ID INNER JOIN PRODUCTS.ID ON USERS.ID=PRODUCTS.ID WHERE USER.ID=%d",u.Id)
	// statement1 := fmt.Sprintf("SELECT * FROM USERS WHERE ID=%d", u.Id)
	// err := db.QueryRow(statement1).Scan(&u.Id, &u.Name, &u.Age)
	// if err != nil {
	// 	return err
	// }
	// statement2 := fmt.Sprintf("SELECT products.id, products.name, products.price FROM products INNER JOIN users ON USERS.ID=PRODUCTS.USER_ID WHERE users.id=%d", u.Id)
	// rows, err := db.Query(statement2)
	// for rows.Next() {
	// 	var product Product
	// 	err := rows.Scan(&product.Id, &product.Name, &product.Price)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	u.Products = append(u.Products, product)
	// }

	
	// m := f.(map[string]interface{})
	// var products []Product

	// res := db.Find(&u)
	// if len(res.GetErrors())>0 {
	// 	return res.GetErrors()
	// }
	// statement := fmt.Sprintf("USER_ID = %d",u.Id)
	// resProduct := db.Where(statement).Find(&u.Products)
	res := db.Model(&u).Related(&u.Products)

	// res := db.Model(&u).Find(products)
	// fmt.Println(res.GetErrors())
	return res.GetErrors()
}

package controller

import (
	"fmt"
	// "errors"
	"github.com/buychat/server/model"
	// "database/sql"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"github.com/buychat/server/respond"
	"github.com/buychat/server/database"
	"strconv"
)

var arr_string_err []string

func AddProduct(w http.ResponseWriter, r *http.Request){
	arr_string_err = arr_string_err[:0]
	decoder := json.NewDecoder(r.Body)
	var product model.Product
	err := decoder.Decode(&product)
	
	if err!= nil {
		arr_string_err = append(arr_string_err, "Invalid Payload Request")
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
		return
	}
	fail := product.AddProduct(database.DB)
	if fail {
		arr_string_err = append(arr_string_err, "Fail Add New Product")
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
		return
	}
	respond.RespondWithJSON(w,http.StatusOK, product)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request){
	arr_string_err = arr_string_err[:0]
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["product_id"])
	if err!= nil {
		arr_string_err = append(arr_string_err, "Invalid Product Id")
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
	}

	p := model.Product{Id:id}
	// var errors []error
	res, errors := p.DeleteProduct(database.DB)
	// fmt.Println(res)
	if len(errors)>0 {
		for _, err  := range errors {
			arr_string_err = append(arr_string_err, err.Error())
		}
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
		return
	}
	respond.RespondWithJSON(w, http.StatusOK, res)
	// if err{
	// 	switch err{
	// 	case sql.ErrNoRows:
	// 		respond.RespondWithError(w, http.StatusBadRequest, "Product Not Found")
	// 	default:
	// 		respond.RespondWithError(w, http.StatusBadRequest, err.Error())
	// 	}
	// 	return
	// }
	// var make(map[]string interface{})
	// respond.RespondWithJSON(w, http.StatusOK, "Delete Success")
	
}

func UpdateProduct(w http.ResponseWriter, r *http.Request){
	arr_string_err = arr_string_err[:0]
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["product_id"])
	if err!= nil {
		arr_string_err = append(arr_string_err,"Invalid Payload")
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
		return
	}
	p := model.Product{Id:id}
	errors := p.UpdateProduct(database.DB)
	if len(errors)>0 {
		for _,err := range errors{
			arr_string_err = append(arr_string_err,err.Error())
		}
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
		return
	}
	respond.RespondWithJSON(w, http.StatusOK, p)
}

func GetAllProduct(w http.ResponseWriter, r *http.Request){
	arr_string_err = arr_string_err[:0]
	query_params := r.URL.Query()
	product_category_id := 0
	if len(query_params) > 0 {
		fmt.Println("dalam")
		id, error := strconv.Atoi(query_params.Get("product_category_id"))
		if error != nil {
			id = 0
		}
		product_category_id = id
	}
	fmt.Println("asdadad")
	products, errors := model.GetAllProduct(database.DB, product_category_id)
	if len(errors)>0 {
		for _,err := range errors{
			arr_string_err = append(arr_string_err, err.Error())
		}
		fmt.Println("dalam error")
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
		return
	}
	fmt.Println("luar error")
	response := map[string]interface{} {
		"count": len(products),
		"data": products,
	}
	respond.RespondWithJSON(w, http.StatusOK, response)
}

func GetProduct(w http.ResponseWriter, r *http.Request){
	arr_string_err = arr_string_err[:0]
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err!= nil {
		arr_string_err = append(arr_string_err,err.Error())
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
		return
	}
	p := model.Product{Id:id}
	errors := p.GetProduct(database.DB)
	if len(errors)>0 {
		for _,err := range errors{
			arr_string_err = append(arr_string_err,err.Error())
		}
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
		return
	}
	respond.RespondWithJSON(w, http.StatusOK, p)
	// if err!= nil {
	// 	switch err{
	// 	case sql.ErrNoRows:
	// 		respond.RespondWithError(w, http.StatusBadRequest, "Product Not Found")
	// 	default:
	// 		respond.RespondWithError(w, http.StatusBadRequest, err.Error())
	// 	}
	// 	return
	// }
	// respond.RespondWithJSON(w, http.StatusOK, p)

}
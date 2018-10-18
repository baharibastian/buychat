package controller

import (
	"github.com/buychat/server/model"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"github.com/buychat/server/respond"
	"github.com/buychat/server/database"
	"strconv"
)

func AddProductCategory(w http.ResponseWriter, r *http.Request) {
	arr_string_err = arr_string_err[:0]
	decoder := json.NewDecoder(r.Body)
	var product_category model.Product_category
	err := decoder.Decode(&product_category)

	if err != nil {
		arr_string_err = append(arr_string_err, "Invalid Payload Request")
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
		return
	}
	fail := product_category.AddCategory(database.DB)
	if fail {
		arr_string_err = append(arr_string_err, "Fail Add New Product Category")
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
		return
	}
	respond.RespondWithJSON(w, http.StatusOK, product_category)
}

func DeleteProductCategory(w http.ResponseWriter, r *http.Request) {
	arr_string_err = arr_string_err[:0]
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		arr_string_err = append(arr_string_err, "Invalid Payload Request")
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
	}

	p := model.Product_category{Id:id}
	res, errors := p.DeleteCategory(database.DB)
	if len(errors) > 0 {
		for _, err := range errors {
			arr_string_err = append(arr_string_err, err.Error())
		}
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
		return
	}
	respond.RespondWithJSON(w, http.StatusOK, res)
}

func GetAllProductCategory(w http.ResponseWriter, r *http.Request){
	arr_string_err = arr_string_err[:0]
	product_categories, errors := model.GetAllProductCategory(database.DB)	
	if len(errors)>0 {
		for _,err := range errors{
			arr_string_err = append(arr_string_err, err.Error())
		}
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
		return
	}
	response := map[string]interface{} {
		"count": len(product_categories),
		"data": product_categories,
	}
	respond.RespondWithJSON(w, http.StatusOK, response)
}
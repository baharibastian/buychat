package controller

import (
	"github.com/buychat/server/model"
	"net/http"
	"vendor/github.com/gorilla/mux"
	"encoding/json"
	"github.com/buychat/server/respond"
	"github.com/buychat/server/database"
	"strconv"
)

func AddDetOrder(w http.ResponseWriter, r *http.Request) {
	arr_string_err = arr_string_err[:0]
	decoder := json.NewDecoder(r.Body)
	var det_order_product model.Det_order_product
	err := decoder.Decode(&det_order_product)

	if err != nil {
		arr_string_err = append(arr_string_err, "Invalid Payload Request")
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
		return
	}
	errors := det_order_product.AddDetOrder(database.DB)
	if len(errors)>0 {
		for _,err := range errors{
			arr_string_err = append(arr_string_err, err.Error())
		}
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
		return
	}
	respond.RespondWithJSON(w, http.StatusOK, det_order_product)
}

func GetDetOrder(w http.ResponseWriter, r *http.Request){
	arr_string_err = arr_string_err[:0]
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err!= nil {
		arr_string_err = append(arr_string_err, "Invalid Detail Order Id")
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
		return
	}
	decoder := json.NewDecoder(r.Body)
	var det_order_product model.Det_order_product
	err = decoder.Decode(&det_order_product)
	if err != nil {
		arr_string_err = append(arr_string_err, "Invalid Payload Request")
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
		return
	}
	det_order_product.Id = id
	errors := det_order_product.AddDetOrder(database.DB)
	if len(errors)>0 {
		for _,err := range errors{
			arr_string_err = append(arr_string_err, err.Error())
		}
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
		return
	}
	respond.RespondWithJSON(w, http.StatusOK, det_order_product)
}

func DeleteDetOrder(w http.ResponseWriter, r *http.Request){
	arr_string_err = arr_string_err[:0]
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err!= nil {
		arr_string_err = append(arr_string_err, "Invalid Detail Order Id")
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
		return
	}
	decoder := json.NewDecoder(r.Body)
	var det_order_product model.Det_order_product
	err = decoder.Decode(&det_order_product)
	if err != nil {
		arr_string_err = append(arr_string_err, "Invalid Payload Request")
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
		return
	}
	det_order_product.Id = id
	errors := det_order_product.DeleteDetOrder(database.DB)
	if len(errors)>0 {
		for _,err := range errors{
			arr_string_err = append(arr_string_err, err.Error())
		}
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
		return
	}
	respond.RespondWithJSON(w, http.StatusOK, det_order_product)
}

func UpdateDetOrder(w http.ResponseWriter, r *http.Request){
	arr_string_err = arr_string_err[:0]
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err!= nil {
		arr_string_err = append(arr_string_err, "Invalid Detail Order Id")
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
		return
	}
	decoder := json.NewDecoder(r.Body)
	var det_order_product model.Det_order_product
	err = decoder.Decode(&det_order_product)
	if err != nil {
		arr_string_err = append(arr_string_err, "Invalid Payload Request")
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
		return
	}
	det_order_product.Id = id
	errors := det_order_product.UpdateDetOrder(database.DB)
	if len(errors)>0 {
		for _,err := range errors{
			arr_string_err = append(arr_string_err, err.Error())
		}
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
		return
	}
	respond.RespondWithJSON(w, http.StatusOK, det_order_product)
}

func GetAllDetOrder(w http.ResponseWriter, r *http.Request){
	arr_string_err = arr_string_err[:0]
	var det_order_products []model.Det_order_product
	det_order_products, errors := model.GetAllDetOrder(database.DB)
	if len(errors)>0 {
		for _,err := range errors{
			arr_string_err = append(arr_string_err, err.Error())
		}
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
		return
	}
	respond.RespondWithJSON(w, http.StatusOK, det_order_products)
}
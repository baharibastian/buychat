package controller

import (
	// "fmt"
	// "errors"
	"buychat/server/model"
	// "database/sql"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"buychat/server/respond"
	"buychat/server/database"
	"strconv"
)

// var arr_string_err []string

func AddOrder(w http.ResponseWriter, r *http.Request){
	arr_string_err = arr_string_err[:0]
	var order model.Order
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&order)
	if err!= nil {
		arr_string_err = append(arr_string_err, "Invalid Payload Request")
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
		return 
	}
	errors := order.AddOrder(database.DB)
	if len(errors)>0 {
		for _,err := range errors{
			arr_string_err = append(arr_string_err, err.Error())
		}
	}
	respond.RespondWithJSON(w, http.StatusOK, order)
}

func DeleteOrder(w http.ResponseWriter, r *http.Request){
	arr_string_err = arr_string_err[:0]
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err!= nil {
		arr_string_err = append(arr_string_err, "Invalid Order Id")
		return
	}
	o := model.Order{Id:id}
	errors := o.DeleteOrder(database.DB)

	if len(errors)>0 {
		for _, err := range errors{
			arr_string_err = append(arr_string_err, err.Error())
		}
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
		return
	}
	respond.RespondWithJSON(w, http.StatusOK, o)
}

func UpdateORder(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err!= nil {
		arr_string_err = append(arr_string_err, "Invalid Order Id")
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
		return
	}
	o := model.Order{Id:id}
	errors := o.UpdateOrder(database.DB)
	if len(errors)>0 {
		for _, err := range errors{
			arr_string_err = append(arr_string_err, err.Error())
		}
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
		return
	}
	respond.RespondWithJSON(w, http.StatusOK, o)
}

func GetOrder(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err!= nil {
		arr_string_err = append(arr_string_err, "Invalid Order Id")
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
		return
	}
	o := model.Order{Id:id}
	errors := o.GetOrder(database.DB)
	if len(errors)>0 {
		for _, err := range errors{
			arr_string_err = append(arr_string_err, err.Error())
		}
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
		return
	}
	respond.RespondWithJSON(w, http.StatusOK, o)	
}

func GetAllOrder(w http.ResponseWriter, r *http.Request){
	var orders []model.Order
	orders, errors := model.GetAllOrder(database.DB)
	if len(errors)>0 {
		for _, err := range errors{
			arr_string_err = append(arr_string_err, err.Error())
		}
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
		return
	}
	respond.RespondWithJSON(w, http.StatusOK, orders)	
}

package controller
 import (
	"buychat/server/model"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"buychat/server/respond"
	"buychat/server/database"
	"strconv"
)

 func AddMerchant(w http.ResponseWriter, r *http.Request) {
	arr_string_err = arr_string_err[:0]
	decoder := json.NewDecoder(r.Body)
	var merchant model.Merchant
	err := decoder.Decode(&merchant)
 	if err!= nil {
		arr_string_err = append(arr_string_err, "Invalid Payload Request")
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
		return
	}
	errors := merchant.AddMerchant(database.DB)
	if len(errors)>0 {
		for _,err := range errors{
			arr_string_err = append(arr_string_err, err.Error())
		}
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
		return
	}
	respond.RespondWithJSON(w,http.StatusOK, merchant)
}

 func GetAllMerchants(w http.ResponseWriter, r *http.Request){
	arr_string_err = arr_string_err[:0]
	merchants, errors := model.GetAllMerchants(database.DB)	
	if len(errors)>0 {
		for _,err := range errors{
			arr_string_err = append(arr_string_err, err.Error())
		}
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
		return
	}
	respond.RespondWithJSON(w, http.StatusOK, merchants)
}

 func GetMerchant(w http.ResponseWriter, r *http.Request){
	arr_string_err = arr_string_err[:0]
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["merchant_id"])
	if err!= nil {
		arr_string_err = append(arr_string_err,err.Error())
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
		return
	}
	p := model.Merchant{Id:id}
	errors := p.GetMerchant(database.DB)
	if len(errors)>0 {
		for _,err := range errors{
			arr_string_err = append(arr_string_err,err.Error())
		}
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
		return
	}
	respond.RespondWithJSON(w, http.StatusOK, p)
} 

func DeleteMerchant(w http.ResponseWriter, r *http.Request){
	arr_string_err = arr_string_err[:0]
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err!= nil {
		arr_string_err = append(arr_string_err, "Invalid Merchant Id")
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
		return
	}
	var m model.Merchant
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&m)
	if err!= nil {
		arr_string_err = append(arr_string_err, "Invalid Payload Request")
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
		return
	}
	m.Id = id
	payload, errors := m.DeleteMerchant(database.DB)
	if len(errors)>0 {
		for _,err := range errors{
			arr_string_err = append(arr_string_err,err.Error())
		}
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
		return
	}
	respond.RespondWithJSON(w, http.StatusOK, payload)
}

func UpdateMerchant(w http.ResponseWriter, r *http.Request){
	arr_string_err = arr_string_err[:0]
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err!= nil {
		arr_string_err = append(arr_string_err, "Invalid Merchant Id")
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
		return
	}
	var m model.Merchant
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&m)
	if err!= nil {
		arr_string_err = append(arr_string_err, "Invalid Payload Request")
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
		return
	}
	m.Id = id
	errors := m.UpdateMerchant(database.DB)
	if len(errors)>0 {
		for _,err := range errors{
			arr_string_err = append(arr_string_err,err.Error())
		}
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
		return
	}
	respond.RespondWithJSON(w, http.StatusOK, m)
}
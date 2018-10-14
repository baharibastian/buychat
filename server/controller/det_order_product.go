package controller

import (
	"buychat/server/model"
	"net/http"
	// "github.com/gorilla/mux"
	"encoding/json"
	"buychat/server/respond"
	"buychat/server/database"
	// "strconv"
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
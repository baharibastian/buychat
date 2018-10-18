package controller

import (
	// "database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"buychat/server/database"
	"buychat/server/model"
	"buychat/server/respond"
	"strconv"

	"github.com/gorilla/mux"
)

func ReadUser(w http.ResponseWriter, r *http.Request) {
	// read params
	arr_string_err = arr_string_err[:0]
	vars := mux.Vars(r)
	// convert string to integer
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal(err)
		arr_string_err = append(arr_string_err, "Invalid User Id")
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
		return
	}
	u := model.User{Id: id}
	new_user, errors := u.GetUser(database.DB)
	if len(errors) > 0 {
		for _, err := range errors {
			arr_string_err = append(arr_string_err, err.Error())
		}
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
		return
	}
	respond.RespondWithJSON(w, http.StatusOK, new_user)
	return
}

func ReadUsers(w http.ResponseWriter, r *http.Request) {
	arr_string_err = arr_string_err[:0]
	users, errors := model.GetUsers(database.DB)
	if len(errors) > 0 {
		for _, err := range errors {
			arr_string_err = append(arr_string_err, err.Error())
		}
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
		return
	}
	respond.RespondWithJSON(w, http.StatusOK, users)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		arr_string_err = append(arr_string_err, "Invalid User Id")
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
		return
	}
	u := model.User{Id: id}
	errors := u.DeleteUser(database.DB)
	if len(errors) > 0 {
		// switch err {
		// case sql.ErrNoRows:
		// 	respond.RespondWithError(w, http.StatusBadRequest, "User Not Found")
		// default:
		// 	respond.RespondWithError(w, http.StatusBadRequest, err.Error())
		// }
		for _, err := range errors {
			arr_string_err = append(arr_string_err, err.Error())
		}
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
		return
	}
	respond.RespondWithJSON(w, http.StatusOK, "User deleted")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		arr_string_err = append(arr_string_err, "Invalid User Id")
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
	}

	var u model.User
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&u)
	if err != nil {
		arr_string_err = append(arr_string_err, "Invalid Request Payload")
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
		return
	}
	defer r.Body.Close()
	u.Id = id
	errors := u.UpdateUser(database.DB)
	if len(errors) > 0 {
		// fmt.Println("error not nil")
		// switch err {
		// case sql.ErrNoRows:
		// 	respond.RespondWithError(w, http.StatusNotFound, "User Not Found")
		// default:
		// 	respond.RespondWithError(w, http.StatusInternalServerError, err.Error())
		// }
		for _, err := range errors {
			arr_string_err = append(arr_string_err, err.Error())
		}
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
		return
	}
	respond.RespondWithJSON(w, http.StatusOK, u)

	// Kalo pake form
	// r.ParseForm()
	// for key, value := range r.Form {
	// 	fmt.Printf("%s = %s\n", key, value)
	// }
	// fmt.Printf("Name => %s\n", r.FormValue("Name"))
	// fmt.Printf("Age => %s\n", r.FormValue("Age"))
	// // age := r.FormValue("Age")
	// var u User
	// u.Id = id
	// u.Name = r.FormValue("Name")
	// u.Age, err = strconv.Atoi(r.FormValue("Age"))
	// if err != nil {
	// 	RespondWithError(w, http.StatusBadRequest, "Invalid Age Person")
	// 	return
	// }
	// err = u.UpdateUser(database.DB)
	// if err != nil {
	// 	switch err {
	// 	case sql.ErrNoRows:
	// 		RespondWithError(w, http.StatusNotFound, "User Not Found")
	// 	default:
	// 		RespondWithError(w, http.StatusBadRequest, err.Error())
	// 		return
	// 	}
	// }
	// RespondWithJSON(w, http.StatusOK, u)
	//

}

func AddUser(w http.ResponseWriter, r *http.Request) {
	var u model.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&u)
	if err != nil {
		arr_string_err = append(arr_string_err, err.Error())
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
	}
	defer r.Body.Close()
	errors := u.AddUser(database.DB)
	if len(errors) > 0 {
		for _, err := range errors {
			arr_string_err = append(arr_string_err, err.Error())
		}
		// arr_string_err = append(arr_string_err, err.Error())
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
	}
	respond.RespondWithJSON(w, http.StatusOK, u)
}

func GetUserProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("User Product")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		arr_string_err = append(arr_string_err, "Invalid User Id")
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
		return
	}
	u := model.User{Id: id}
	fmt.Println("get users")
	errors := u.GetUserProduct(database.DB)
	fmt.Println("get products")
	if len(errors) > 0 {
		// switch err {
		// case sql.ErrNoRows:
		// 	respond.RespondWithError(w, http.StatusBadRequest, "User Notr Found")
		// default:
		// 	respond.RespondWithError(w, http.StatusBadRequest, err.Error())
		// }
		for _, err := range errors {
			arr_string_err = append(arr_string_err, err.Error())
		}
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
		return
	}
	respond.RespondWithJSON(w, http.StatusOK, u)

}

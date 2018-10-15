package route

import (
	"buychat/server/controller"

	"github.com/gorilla/mux"
)

var Router *mux.Router

func init() {
	Router = mux.NewRouter()

	Router.HandleFunc("/user/product/{id:[0-9]+}", controller.GetUserProduct).Methods("GET")
	Router.HandleFunc("/user/{id:[0-9]+}", controller.ReadUser).Methods("GET")
	Router.HandleFunc("/user/{id:[0-9]+}", controller.UpdateUser).Methods("PUT")
	Router.HandleFunc("/users", controller.ReadUsers).Methods("GET")
	Router.HandleFunc("/users/{id:[0-9]}", controller.DeleteUser).Methods("DELETE")
	Router.HandleFunc("/users", controller.AddUser).Methods("POST")
	Router.HandleFunc("/products/{id:[0-9]+}", controller.GetProduct).Methods("GET")
	Router.HandleFunc("/products/{id:[0-9]+}", controller.UpdateProduct).Methods("PUT")
	Router.HandleFunc("/products", controller.GetAllProduct).Methods("GET")
	Router.HandleFunc("/products/{id:[0-9]+}", controller.DeleteProduct).Methods("DELETE")
	Router.HandleFunc("/products", controller.AddProduct).Methods("POST")
	Router.HandleFunc("/orders", controller.AddOrder).Methods("POST")
	Router.HandleFunc("/orders", controller.GetAllOrder).Methods("GET")
	Router.HandleFunc("/orders/{id:[0-9]+}", controller.UpdateOrder).Methods("PUT")
	Router.HandleFunc("/orders/{id:[0-9]+}", controller.DeleteOrder).Methods("DELETE")
	Router.HandleFunc("/orders/{id:[0-9]+}", controller.GetOrder).Methods("GET")
	Router.HandleFunc("/merchant", controller.AddMerchant).Methods("POST")
	Router.HandleFunc("/merchant/{id:[0-9]+}", controller.DeleteMerchant).Methods("DELETE")
	Router.HandleFunc("/merchant", controller.GetAllMerchants).Methods("GET")

}

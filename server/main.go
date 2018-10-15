package main

import (
	"log"
	"net/http"
	"buychat/server/route"
)

func main() {

	log.Fatal(http.ListenAndServe(":8000", route.Router))
}
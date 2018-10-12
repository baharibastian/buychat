package main

import (
	"log"
	"net/http"
	"rest_api/route"
)

func main() {

	log.Fatal(http.ListenAndServe(":8000", route.Router))
}

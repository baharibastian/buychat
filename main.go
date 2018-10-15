package main

import (
	"log"
	"net/http"
	"buychat/server/route"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = ""
	}
	log.Fatal(http.ListenAndServe(":"+port, route.Router))
}
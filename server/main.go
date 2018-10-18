package main

import (
	"log"
	"net/http"
	"github.com/buychat/server/route"
	"os"
	"fmt"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = ""
	}
	log.Fatal(http.ListenAndServe(":8000", route.Router))
	fmt.Println("REST server run on localhost:" + port)
}
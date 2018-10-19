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
		port = "8000"
	}
	log.Fatal(http.ListenAndServe(":"+port, route.Router))
	fmt.Println("REST server run on localhost:" + port)
}
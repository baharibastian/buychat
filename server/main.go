package main

import (
	"log"
	"net/http"
<<<<<<< HEAD
	"buychat/server/route"
	"os"
=======
	"github.com/buychat/server/route"
	"os"
	"fmt"
>>>>>>> 70a94b4de082a58cc8f742cb6dd36f725068e634
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
<<<<<<< HEAD
		port = ""
	}
	log.Fatal(http.ListenAndServe(":"+port, route.Router))
=======
		port = "8000"
	}
	log.Fatal(http.ListenAndServe(":"+port, route.Router))
	fmt.Println("REST server run on localhost:" + port)
>>>>>>> 70a94b4de082a58cc8f742cb6dd36f725068e634
}
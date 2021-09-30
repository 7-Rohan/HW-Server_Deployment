package main

import (
	"log"
	"net/http"
	"os"

	"github.com/7-Rohan/HW-Server_Deployment/router"
)

func main() {
	port := os.Getenv("PORT")
	http.Handle("/", router.Router())
	log.Print("Listening on:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

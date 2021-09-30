package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type JsonUserid struct {
	Username string `json:"username"`
	Follower int    `json:"followers"`
}

func home(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Welcome")
}

func sammy(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	jsonUserUserid := vars["userid"]

	jsonUser := map[string]JsonUserid{}

	jsonFile, err := os.Open("json-HW.json")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &jsonUser)

	result := jsonUser[jsonUserUserid]
	json.NewEncoder(w).Encode(result)
}

func main() {
	r := mux.NewRouter()
	port := os.Getenv("PORT")
	r.HandleFunc("/", home)
	r.HandleFunc("/{userid}", sammy)
	http.Handle("/", r)
	log.Print("Listening on:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

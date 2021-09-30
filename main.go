package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type JsonUserid struct {
	Username string `json:"username"`
	Follower int    `json:"followers"`
}

func home(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Welcome")
}

func sammy(w http.ResponseWriter, r *http.Request) {
	jsonUser := map[string]JsonUserid{}

	jsonFile, err := os.Open("json-HW.json")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &jsonUser)

	result := jsonUser["sammy"]
	json.NewEncoder(w).Encode(result)
}

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", home)
	http.HandleFunc("/sammy", sammy)
	log.Print("Listening on:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

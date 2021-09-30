package middleware

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/7-Rohan/HW-Server_Deployment/model"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, " Welcome to Homework-Server_Deployment Rakamin \n\n\n\n Untuk mengakses salah satu data user: https://hwheroku-deployment-18797.herokuapp.com/{userid}, contohnya https://hwheroku-deployment-18797.herokuapp.com/sammy \n\n Untuk mengakses jumlah follower yang dimiliki suatu user: https://hwheroku-deployment-18797.herokuapp.com/follower/{username}, contohnya https://hwheroku-deployment-18797.herokuapp.com/follower/SammyShark \n\n Terima kasih \n\n Berikut list data user yang ada: \n\n")

	jsonUsers := map[string]model.JsonUserid{}

	jsonFile, err := os.Open("json-HW.json")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &jsonUsers)
	defer jsonFile.Close()
	json.NewEncoder(w).Encode(jsonUsers)
}

func GetByUserid(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	jsonUserUserid := vars["userid"]

	jsonUser := map[string]model.JsonUserid{}

	jsonFile, err := os.Open("json-HW.json")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &jsonUser)

	result, ok := jsonUser[jsonUserUserid]
	if !ok {
		json.NewEncoder(w).Encode(fmt.Sprintf("User %s tidak tersedia di file json-HW.json", jsonUserUserid))
	} else {
		json.NewEncoder(w).Encode(result)
	}
	defer jsonFile.Close()
}

func GetByUsername(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	jsonUserUsername := vars["username"]

	jsonUsers := map[string]model.JsonUserid{}

	jsonFile, err := os.Open("json-HW.json")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &jsonUsers)

	var result model.JsonUserid

	for _, v := range jsonUsers {
		if v.Username == jsonUserUsername {
			result = v
		}
	}

	if result.Follower == 0 {
		json.NewEncoder(w).Encode(fmt.Sprintf("Username %s tidak tersedia di file json-HW.json", jsonUserUsername))
	} else {
		str := []string{"followers: ", strconv.Itoa(result.Follower)}
		res := strings.Join(str, " ")

		json.NewEncoder(w).Encode(res)
	}

	defer jsonFile.Close()
}

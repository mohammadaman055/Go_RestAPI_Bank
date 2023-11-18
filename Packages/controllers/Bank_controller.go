package controllers

import (
	"encoding/json"
	"fmt"
	"github/mohammadaman055/Go_RestAPI_Bank/Packages/modules"
	"github/mohammadaman055/Go_RestAPI_Bank/Packages/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var NewUser modules.User

func GetUsers(w http.ResponseWriter, r *http.Request) {
	newusers := modules.GetUsers()
	result, _ := json.Marshal(newusers)
	w.Header().Set("Content-TYpe", "application/json")
	w.WriteHeader((http.StatusOK))
	w.Write(result)
}

func GetUserBYid(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userid := vars["userid"]
	ID, err := strconv.ParseInt(userid, 0, 0)
	if err != nil {
		fmt.Println("error on parsing")
	}
	userdetail := modules.GetUserBYid(ID)
	result, _ := json.Marshal(userdetail)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func NewUsers(w http.ResponseWriter, r *http.Request) {
	CreateUser := &modules.User{}
	utils.ParseData(r, CreateUser)
	fmt.Printf("Parsed data: %+v\n", CreateUser)
	u := CreateUser.NewUsers()
	result, _ := json.Marshal(u)
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userid := vars["userid"]
	ID, err := strconv.ParseInt(userid, 0, 0)
	if err != nil {
		fmt.Println("error on parsing")
	}
	deleteuserdetail := modules.DeleteUser(ID)
	result, _ := json.Marshal(deleteuserdetail)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
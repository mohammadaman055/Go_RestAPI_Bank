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

func GetAllAccount(w http.ResponseWriter, r *http.Request) {
	newuaccount := modules.GetAllAccount()
	result, _ := json.Marshal(newuaccount)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader((http.StatusOK))
	w.Write(result)
}

func GetAccById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accid := vars["accid"]
	ID, err := strconv.ParseInt(accid, 0, 0)
	if err != nil {
		fmt.Println("error on parsing")
	}
	accdetail := modules.GetAccById(ID)
	result, _ := json.Marshal(accdetail)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func NewAccount(w http.ResponseWriter, r *http.Request) {
	CreateAcc := &modules.Account{}
	utils.ParseData(r, CreateAcc)
	fmt.Printf("Parsed data: %+v\n", CreateAcc)
	a := CreateAcc.NewAccount()
	result, _ := json.Marshal(a)
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func DeleteAccById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accid := vars["accid"]
	ID, err := strconv.ParseInt(accid, 0, 0)
	if err != nil {
		fmt.Println("error on parsing")
	}
	modules.DeleteAccById(ID)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Account deleted successfully"))
}

func UpdateAccById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accid := vars["accid"]

	ID, err := strconv.ParseInt(accid, 0, 0)
	if err != nil {
		fmt.Println("error on parsing ID")
	}

	updateAcc := &modules.Account{}
	utils.ParseData(r, updateAcc)

	updateAcc.Accountid = int(ID)
	updateAcc.UpdateAccById()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Account updated successfully"))
}

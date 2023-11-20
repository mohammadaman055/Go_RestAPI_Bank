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

func AllTransactions(w http.ResponseWriter, r *http.Request) {
	transactiondata := modules.AllTransactions()
	result, _ := json.Marshal(transactiondata)

	w.Header().Set("Content-TYpe", "application/json")
	w.WriteHeader((http.StatusOK))
	w.Write(result)
}

func NewTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accid := vars["accid"]

	ID, err := strconv.ParseInt(accid, 0, 0)
	if err != nil {
		fmt.Println("error on parsing ID")
	}

	transaction := &modules.Transaction{}
	utils.ParseData(r, transaction)

	transaction.Accountid = int(ID)
	transaction.NewTransaction()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Transaction successfully"))
}

func Statement(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accid := vars["accid"]

	ID, err := strconv.ParseInt(accid, 0, 0)
	if err != nil {
		fmt.Println("error on parsing ID")
	}
	statdetail := modules.Statement(ID)
	result, _ := json.Marshal(statdetail)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

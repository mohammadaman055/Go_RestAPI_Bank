package main

import (
	"fmt"
	"log"
	"net/http"

	"github/mohammadaman055/Go_RestAPI_Bank/Packages/routes"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.BankRoutes(r)

	http.Handle("/", r)
	log.Fatal((http.ListenAndServe(":8080", r)))
	fmt.Println("Server Running at localhost:8080")
}

package routes

import (
	"github/mohammadaman055/Go_RestAPI_Bank/Packages/controllers"

	"github.com/gorilla/mux"
)

var BankRoutes = func(router *mux.Router) {

	router.HandleFunc("/users/", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/user/", controllers.NewUsers).Methods("POST")
	router.HandleFunc("/user/{userid}", controllers.GetUserBYid).Methods("GET")
	router.HandleFunc("/user/{userid}", controllers.UpdateUserBYid).Methods("PUT")
	router.HandleFunc("/user/{userid}", controllers.DeleteUser).Methods("DELETE")

	router.HandleFunc("/accounts/", controllers.GetAllAccount).Methods("GET")
	router.HandleFunc("/account/", controllers.NewAccount).Methods("POST")
	router.HandleFunc("/account/{accid}", controllers.GetAccById).Methods("GET")
	router.HandleFunc("/account/{accid}", controllers.UpdateAccById).Methods("PUT")
	router.HandleFunc("/account/{accid}", controllers.DeleteAccById).Methods("DELETE")

	router.HandleFunc("/transactions/", controllers.AllTransactions).Methods("GET")
	router.HandleFunc("/transaction/{accid}", controllers.NewTransaction).Methods("POST")
	router.HandleFunc("/statement/{accid}", controllers.Statement).Methods("GET")
}

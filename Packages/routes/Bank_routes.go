package routes

import (
	"github/mohammadaman055/Go_RestAPI_Bank/Packages/controllers"

	"github.com/gorilla/mux"
)

var BankRoutes = func(router *mux.Router) {

	router.HandleFunc("/users/", controllers.GetUsers).Methods("GET")              //extract all customer data
	router.HandleFunc("/user/", controllers.NewUsers).Methods("POST")              //add or create new usres
	router.HandleFunc("/user/{userid}", controllers.GetUserBYid).Methods("GET")    //one user data by ID
	router.HandleFunc("/user/{userid}", controllers.UpdateUserBYid).Methods("PUT") //update one user by ID
	router.HandleFunc("/user/{userid}", controllers.DeleteUser).Methods("DELETE")  //Delete user by id

	router.HandleFunc("/accounts/", controllers.GetAllAccount).Methods("GEt")
	router.HandleFunc("/account/", controllers.NewAccount).Methods("POST")
	router.HandleFunc("/account/{accid}", controllers.GetAccById).Methods("GET")
	router.HandleFunc("/account/{accid}", controllers.UpdateAccById).Methods("PUT")
	router.HandleFunc("/account/{accid}", controllers.DeleteAccById).Methods("DELETE")
}

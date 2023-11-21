package main

import "fmt"

func main() {
<<<<<<< HEAD
	r := mux.NewRouter()
	routes.BankRoutes(r)

	http.Handle("/", r)
	fmt.Println("Server Running at localhost:8080")
	log.Fatal((http.ListenAndServe(":8080", r)))
=======
	fmt.Println("Hi Remote Repo ON!")
>>>>>>> parent of d826c8c (CRD API Complete)
}

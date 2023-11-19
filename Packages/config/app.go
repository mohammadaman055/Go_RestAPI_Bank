package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Connect() {
	d, err := sql.Open("mysql", "root:Aman@2002@tcp(127.0.0.1:3306)/bankapi")
	if err != nil {
		log.Fatal(err)
	}
	// defer d.Close()

	err = d.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Connected to the Mysql database ....")

	db = d
}

func GetDB() *sql.DB {
	return db
}

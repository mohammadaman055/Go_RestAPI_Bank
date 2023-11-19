package modules

import (
	"database/sql"
	"fmt"
	"github/mohammadaman055/Go_RestAPI_Bank/Packages/config"
)

var db *sql.DB

type Users struct {
	UsersID int    `json:"id"`
	Name    string `json:"name"`
	Yob     string `json:"yob"`
	Address string `json:"address"`
	Phone   int    `json:"phone"`
}

func init() {
	config.Connect()
	db = config.GetDB()
}

func (u *Users) NewUsers() *Users {
	_, err := db.Query("INSERT INTO Users (name, yob, address, phone) VALUES (?, ?, ?, ?)", u.Name, u.Yob, u.Address, u.Phone)
	if err != nil {
		panic(err.Error())
	}
	return u
}

func GetUsers() []Users {
	var allUsers []Users

	Usersdata, err := db.Query("SELECT * FROM Users")
	if err != nil {
		fmt.Println("Error on retrieving")
		panic(err.Error())
	}

	for Usersdata.Next() {
		var User Users
		err := Usersdata.Scan(&User.UsersID, &User.Name, &User.Yob, &User.Address, &User.Phone)
		if err != nil {
			panic(err.Error())
		}

		allUsers = append(allUsers, User)
	}

	return allUsers
}

func GetUserBYid(UsersID int64) *Users {
	var getUsers Users

	row := db.QueryRow("SELECT * FROM Users WHERE userid=?", UsersID)

	err := row.Scan(&getUsers.UsersID, &getUsers.Name, &getUsers.Yob, &getUsers.Address, &getUsers.Phone)
	if err != nil {
		fmt.Println("Error retrieving Users:", err)
		return nil
	}

	return &getUsers
}

func DeleteUser(UsersID int64) {
	_, err := db.Exec("DELETE FROM Users WHERE userid=?", UsersID)
	if err != nil {
		panic(err.Error())
	}
}

func (u *Users) UpdateUserBYid() {
	_, err := db.Exec("UPDATE Users SET name=?, yob=?, address=?, phone=? WHERE userid=?", u.Name, u.Yob, u.Address, u.Phone, u.UsersID)
	if err != nil {
		panic(err.Error())
	}
}

package modules

import (
	"database/sql"
	"fmt"
	"github/mohammadaman055/Go_RestAPI_Bank/Packages/config"
)

var db *sql.DB

type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Age     string `json:"age"`
	Phone   int    `json:"phone"`
	Balance int    `json:"balance"`
}

func init() {
	config.Connect()
	db = config.GetDB()
}

func (u *User) NewUsers() *User {
	_, err := db.Query("INSERT INTO users (id,name, age, phone, balance) VALUES (?,?, ?, ?, ?)", u.ID, u.Name, u.Age, u.Phone, u.Balance)
	if err != nil {
		panic(err.Error())
	}
	return u
}

func GetUsers() []User {
	var users []User

	userdata, err := db.Query("SELECT * FROM users")
	if err != nil {
		fmt.Println("Error on retrieving")
		panic(err.Error())
	}
	// defer userdata.Close()

	for userdata.Next() {
		var user User
		err := userdata.Scan(&user.ID, &user.Name, &user.Age, &user.Phone, &user.Balance)
		if err != nil {
			panic(err.Error())
		}

		users = append(users, user)
	}

	return users
}

func GetUserBYid(userID int64) *User {
	var getUser User

	row := db.QueryRow("SELECT * FROM users WHERE id=?", userID)

	err := row.Scan(&getUser.ID, &getUser.Name, &getUser.Age, &getUser.Phone, &getUser.Balance)
	if err != nil {
		fmt.Println("Error retrieving user:", err)
		return nil
	}

	return &getUser
}

func DeleteUser(userID int64) *User {
	var deletedUser User

	row := db.QueryRow("SELECT * FROM users WHERE id=?", userID)

	err := row.Scan(&deletedUser.ID, &deletedUser.Name, &deletedUser.Age, &deletedUser.Phone, &deletedUser.Balance)
	if err != nil {
		fmt.Println("Error retrieving user before deletion:", err)
		panic(err.Error())
	}

	db.Exec("DELETE FROM users WHERE id=?", userID)

	return &deletedUser
}

func (u *User)UpdateUserBYid()  {
	 _, err := db.Exec("UPDATE users SET name=?, age=?, phone=?, balance=? WHERE id=?", u.Name, u.Age, u.Phone, u.Balance, u.ID)
    if err != nil {
        panic(err.Error())
    }
}

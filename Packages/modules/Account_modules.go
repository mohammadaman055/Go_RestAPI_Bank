package modules

import "fmt"

type Account struct {
	Accountid    int    `json:"accountid"`
	Userid       int    `json:"userid"`
	Account_type string `json:"account_type"`
	Balance      int    `json:"balance"`
}

func (a *Account) NewAccount() *Account {
	_, err := db.Query("INSERT INTO accounts (userid, account_type, balance) VALUES (?, ?, ?)", a.Userid, a.Account_type, a.Balance)
	if err != nil {
		panic(err.Error())
	}
	return a
}

func GetAllAccount() []Account {
	var allAccount []Account

	accountdata, err := db.Query("SELECT * FROM accounts")
	if err != nil {
		fmt.Println("Error on retrieving")
		panic(err.Error())
	}

	for accountdata.Next() {
		var account Account
		err := accountdata.Scan(&account.Accountid, &account.Userid, &account.Account_type, &account.Balance)
		if err != nil {
			panic(err.Error())
		}

		allAccount = append(allAccount, account)
	}

	return allAccount
}

func GetAccById(Accid int64) *Account {
	var getAccount Account

	row := db.QueryRow("SELECT * FROM accounts WHERE accountid=?", Accid)

	err := row.Scan(&getAccount.Accountid, &getAccount.Userid, &getAccount.Account_type, &getAccount.Balance)
	if err != nil {
		fmt.Println("Error retrieving Users:", err)
		return nil
	}

	return &getAccount
}

func DeleteAccById(Accid int64) {
	_, err := db.Exec("DELETE FROM accounts WHERE accountid=?", Accid)
	if err != nil {
		panic(err.Error())
	}
}

func (a *Account) UpdateAccById() {
	_, err := db.Exec("UPDATE accounts SET userid=?, account_type=?, balance=? WHERE accountid=?", a.Userid, a.Account_type, a.Balance, a.Accountid)
	if err != nil {
		panic(err.Error())
	}
}

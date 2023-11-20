package modules

import (
	"fmt"
)

type Transaction struct {
	Transaction_id   int    `json:"transaction_id"`
	Accountid        int    `json:"accountid"`
	Recvid           int    `json:"recvid"`
	Amount           int    `json:"amount"`
	Transaction_type string `json:"transaction_type"`
}

func AllTransactions() []Transaction {
	var transactions []Transaction
	tdata, err := db.Query("SELECT * FROM transactions")
	if err != nil {
		fmt.Println("Error on retrieving all transactions")
		panic(err.Error())
	}
	for tdata.Next() {
		var transaction Transaction

		err := tdata.Scan(&transaction.Transaction_id, &transaction.Accountid, &transaction.Recvid, &transaction.Amount, &transaction.Transaction_type)
		if err != nil {
			panic(err.Error())
		}

		transactions = append(transactions, transaction)
	}

	return transactions
}

func (t *Transaction) NewTransaction() {
	if t.Recvid != 0 {
		db.Query("INSERT INTO transactions (accountid, recvid, amount, transaction_type) VALUES (?, ?, ?, ?)", t.Accountid, t.Recvid, t.Amount, t.Transaction_type)

		balanceRow, err := db.Query("SELECT balance FROM accounts WHERE accountid=?", t.Accountid)
		if err != nil {
			panic(err.Error())
		}
		defer balanceRow.Close()

		if balanceRow.Next() {
			var balance int
			if err := balanceRow.Scan(&balance); err != nil {
				panic(err.Error())
			}

			newBalance := balance - t.Amount
			_, err := db.Query("UPDATE accounts SET balance = ? WHERE accountid=?", newBalance, t.Accountid)
			if err != nil {
				panic(err.Error())
			}
		}

		rbalanceRow, err := db.Query("SELECT balance FROM accounts WHERE accountid=?", t.Recvid)
		if err != nil {
			panic(err.Error())
		}
		defer rbalanceRow.Close()

		if rbalanceRow.Next() {
			var rbalance int
			if err := rbalanceRow.Scan(&rbalance); err != nil {
				panic(err.Error())
			}

			newBalance := rbalance + t.Amount
			_, err := db.Query("UPDATE accounts SET balance = ? WHERE accountid=?", newBalance, t.Recvid)
			if err != nil {
				panic(err.Error())
			}
		}
		db.Query("INSERT INTO transactions (accountid, recvid, amount, transaction_type) VALUES (?, ?, ?, 'Credited')", t.Recvid, t.Accountid, t.Amount)

	} else {
		db.Query("INSERT INTO transactions (accountid, recvid, amount, transaction_type) VALUES (?, ?, ?, ?)", t.Accountid, t.Recvid, t.Amount, t.Transaction_type)

		balanceRow, err := db.Query("SELECT balance FROM accounts WHERE accountid=?", t.Accountid)
		if err != nil {
			panic(err.Error())
		}
		defer balanceRow.Close()

		if balanceRow.Next() {
			var balance int
			if err := balanceRow.Scan(&balance); err != nil {
				panic(err.Error())
			}

			newBalance := balance + t.Amount
			_, err := db.Query("UPDATE accounts SET balance = ? WHERE accountid=?", newBalance, t.Accountid)
			if err != nil {
				panic(err.Error())
			}
		}
	}
}

func Statement(accid int64) []Transaction {

	var stat []Transaction

	statdata, err := db.Query("SELECT * FROM transactions WHERE accountid=?", accid)
	if err != nil {
		fmt.Println("Error on retrieving")
		panic(err.Error())
	}

	for statdata.Next() {
		var getstat Transaction
		err := statdata.Scan(&getstat.Transaction_id, &getstat.Accountid, &getstat.Recvid, &getstat.Amount, &getstat.Transaction_type)
		if err != nil {
			panic(err.Error())
		}

		stat = append(stat, getstat)
	}

	return stat

}

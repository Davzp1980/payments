package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func CreatePayment(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input InputPayment
		var reciever User
		var account Account
		var payment Payment

		json.NewDecoder(r.Body).Decode(&input)

		err := db.QueryRow("SELECT id, name FROM users WHERE name=$1", input.PayerName).Scan(&reciever.ID, &reciever.Name)
		if err != nil {
			log.Println("User does not exists")
			w.WriteHeader(http.StatusForbidden)
		}
		fmt.Println(reciever.ID, reciever.Name, input.AmountPayment)

		err = db.QueryRow("SELECT id,balance  FROM accounts WHERE user_id=$1", reciever.ID).Scan(&account.ID, &account.Balance)
		if err != nil {
			log.Println("Account does not exists")
			w.WriteHeader(http.StatusForbidden)
		}
		fmt.Println(account.ID)

		err = db.QueryRow("INSERT INTO payments (account_id, amount_payment, date, reciever) VALUES ($1,$2,$3,$4) RETURNING id",
			account.ID, input.AmountPayment, time.Now(), input.ReceiverName).Scan(
			&payment.ID)
		if err != nil {
			log.Println("Create payment error")
			w.WriteHeader(http.StatusForbidden)
		}

		balance := account.Balance + input.AmountPayment
		_, err = db.Exec("UPDATE accounts SET balance=$2 WHERE id=$1", account.ID, balance)
		if err != nil {
			log.Println("Add to balance error")
			w.WriteHeader(http.StatusForbidden)
		}

	}
}

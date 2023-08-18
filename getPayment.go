package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func GetPaymentsById(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input Input

		var account Account
		var user User

		json.NewDecoder(r.Body).Decode(&input)

		err := db.QueryRow("SELECT id FROM users WHERE name=$1", input.Name).Scan(&user.ID)
		if err != nil {
			log.Println("User does not exists")
			w.WriteHeader(http.StatusNotFound)
		}
		fmt.Println("user.ID", user.ID)

		err = db.QueryRow("SELECT id FROM accounts WHERE user_id=$1", user.ID).Scan(&account.ID)
		if err != nil {
			log.Println("User does not exists")
			w.WriteHeader(http.StatusNotFound)
		}
		fmt.Println("account.ID", account.ID)
		rows, err := db.Query("SELECT * FROM payments where user_id=$1 ORDER BY id", user.ID)
		if err != nil {
			log.Panicln("Account selection error")
			w.WriteHeader(http.StatusNotFound)
		}
		fmt.Println(account.ID)
		sortedPayments := []Payment{}

		for rows.Next() {
			var p Payment

			if err = rows.Scan(&p.ID, &p.UserId, &p.Reciever, &p.RecieverIban, &p.Payer, &p.PayerIban, &p.AmountPayment, &p.Date); err != nil {
				log.Println(err)
			}
			sortedPayments = append(sortedPayments, p)
		}

		json.NewEncoder(w).Encode(sortedPayments)

	}
}

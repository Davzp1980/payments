package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

func GetAccountsById(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		/*var input Input

		var user User

		json.NewDecoder(r.Body).Decode(&input)

		err := db.QueryRow("SELECT id FROM users WHERE name=$1", input.Name).Scan(&user.ID)
		if err != nil {
			log.Println("User does not exists")
			w.WriteHeader(http.StatusNotFound)
		}
		*/
		rows, err := db.Query("SELECT * FROM accounts ORDER BY id")
		if err != nil {
			log.Panicln("Account selection error")
			w.WriteHeader(http.StatusNotFound)
		}

		sortedAccounts := []Account{}

		for rows.Next() {
			var a Account

			if err = rows.Scan(&a.ID, &a.UserId, &a.Iban, &a.Balance); err != nil {
				log.Println(err)
			}
			sortedAccounts = append(sortedAccounts, a)
		}

		json.NewEncoder(w).Encode(sortedAccounts)

	}
}

func GetAccountsByIban(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		/*var input Input

		var user User

		json.NewDecoder(r.Body).Decode(&input)

		err := db.QueryRow("SELECT id FROM users WHERE name=$1", input.Name).Scan(&user.ID)
		if err != nil {
			log.Println("User does not exists")
			w.WriteHeader(http.StatusNotFound)
		}
		*/
		rows, err := db.Query("SELECT * FROM accounts ORDER BY iban")
		if err != nil {
			log.Panicln("Account selection error")
			w.WriteHeader(http.StatusNotFound)
		}
		sortedAccounts := []Account{}

		for rows.Next() {
			var a Account

			if err = rows.Scan(&a.ID, &a.UserId, &a.Iban, &a.Balance); err != nil {
				log.Println(err)
			}
			sortedAccounts = append(sortedAccounts, a)
		}

		json.NewEncoder(w).Encode(sortedAccounts)

	}
}

func GetAccountsByBalance(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		/*var input Input

		var user User

		json.NewDecoder(r.Body).Decode(&input)

		err := db.QueryRow("SELECT id FROM users WHERE name=$1", input.Name).Scan(&user.ID)
		if err != nil {
			log.Println("User does not exists")
			w.WriteHeader(http.StatusNotFound)
		}
		*/
		rows, err := db.Query("SELECT * FROM accounts ORDER BY balance")
		if err != nil {
			log.Panicln("Account selection error")
			w.WriteHeader(http.StatusNotFound)
		}
		sortedAccounts := []Account{}

		for rows.Next() {
			var a Account

			if err = rows.Scan(&a.ID, &a.UserId, &a.Iban, &a.Balance); err != nil {
				log.Println(err)
			}
			sortedAccounts = append(sortedAccounts, a)
		}

		json.NewEncoder(w).Encode(sortedAccounts)

	}
}

package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

func CreateAccount(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input Input
		var user User
		var account Account

		json.NewDecoder(r.Body).Decode(&input)

		err := db.QueryRow("SELECT id FROM users WHERE name=$1", input.Name).Scan(&user.ID)
		if err != nil {
			log.Println("User does not exists")
			w.WriteHeader(http.StatusForbidden)
		}
		iban := "000" + input.Name

		err = db.QueryRow("INSERT INTO accounts (user_id, iban) VALUES ($1,$2) RETURNING id", user.ID, iban).Scan(
			&account.ID, &account.ID, &account.Iban, &account.Balance)
		if err != nil {
			log.Println("Create account error")
			w.WriteHeader(http.StatusBadRequest)
		}

	}
}

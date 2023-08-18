package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
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
			return
		}
		i := strconv.Itoa(rand.Intn(1000000000))
		iban := i + input.Name
		fmt.Println(iban)

		err = db.QueryRow("INSERT INTO accounts (user_id, iban) VALUES ($1,$2) RETURNING id", user.ID, iban).Scan(
			&account.ID)
		if err != nil {
			log.Println("Create account error")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

	}
}

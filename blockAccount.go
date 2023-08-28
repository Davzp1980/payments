package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//Поля в postman:
// "iban"

func BlockAccount(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input Input
		var account Account

		json.NewDecoder(r.Body).Decode(&input)

		err := db.QueryRow("SELECT iban, blocked FROM accounts WHERE iban=$1", input.Iban).Scan(&account.Iban, &account.Blocked)
		if err != nil {
			log.Println("Here", err)
		}
		expectedIban := account.Iban
		inputIban := input.Iban

		if inputIban == expectedIban {
			_, err := db.Exec("UPDATE accounts SET blocked=$1 WHERE iban=$2", true, input.Iban)
			if err != nil {
				log.Println(err)
			}

		} else {
			log.Println("Account", inputIban, "Does not exists")
			return
		}
		w.Write([]byte(fmt.Sprintf("Account %s blocked", input.Name)))

	}
}

func UnBlockAccount(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input Input
		var account Account

		json.NewDecoder(r.Body).Decode(&input)

		err := db.QueryRow("SELECT iban, blocked FROM accounts WHERE iban=$1", input.Iban).Scan(&account.Iban, &account.Blocked)
		if err != nil {
			log.Println("Here", err)
		}
		expectedIban := account.Iban
		inputIban := input.Iban

		if inputIban == expectedIban {
			_, err := db.Exec("UPDATE accounts SET blocked=$1 WHERE iban=$2", false, input.Iban)
			if err != nil {
				log.Println(err)
			}

		} else {
			log.Println("Account", inputIban, "Does not exists")
			return
		}
		w.Write([]byte(fmt.Sprintf("Account %s unblocked", input.Name)))

	}
}

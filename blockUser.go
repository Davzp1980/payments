package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func BlockUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input Input
		var user User

		json.NewDecoder(r.Body).Decode(&input)

		err := db.QueryRow("SELECT name FROM users WHERE name=$1", input.Name).Scan(&user.Name)
		if err != nil {
			log.Println("Here", err)
		}
		expectedName := user.Name
		inputName := input.Name

		if inputName == expectedName {
			_, err := db.Exec("UPDATE users SET blocked=$1 WHERE name=$2", true, input.Name)
			if err != nil {
				log.Println(err)
			}

		} else {
			log.Println("User", inputName, "Does not exists")
			return
		}
		w.Write([]byte(fmt.Sprintf("User %s blocked", input.Name)))

	}
}

func UnBlockUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input Input
		var user User

		json.NewDecoder(r.Body).Decode(&input)

		err := db.QueryRow("SELECT name FROM users WHERE name=$1", input.Name).Scan(&user.Name)
		if err != nil {
			log.Println("Here", err)
		}
		expectedName := user.Name
		inputName := input.Name

		if inputName == expectedName {
			_, err := db.Exec("UPDATE users SET blocked=$1 WHERE name=$2", false, input.Name)
			if err != nil {
				log.Println(err)

			}

		} else {
			log.Println("User", inputName, "Does not exists")
			return
		}
		w.Write([]byte(fmt.Sprintf("User %s unblocked", input.Name)))

	}
}

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func CreateAdmin(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input Input
		var user User

		json.NewDecoder(r.Body).Decode(&input)
		isAdmin := true
		hashedPassword, _ := HashedPassword(input.Password)

		err := db.QueryRow("SELECT name FROM users WHERE name=$1", input.Name).Scan(&user.Name)
		if err != nil {
			log.Println("User does not exixts")
			return
		}

		_, err = db.Query("INSERT INTO users (name, password, is_admin) VALUES ($1,$2,$3)", input.Name, hashedPassword, isAdmin)
		if err != nil {
			log.Println(err)
			return
		}
		w.Write([]byte(fmt.Sprintf("User %s created", input.Name)))
	}

}

func CreateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input Input
		var user User

		json.NewDecoder(r.Body).Decode(&input)
		isAdmin := false
		hashedPassword, _ := HashedPassword(input.Password)

		err := db.QueryRow("SELECT name FROM users WHERE name=$1", input.Name).Scan(&user.Name)
		if err != nil {
			log.Println("User does not exixts")
			return
		}
		expectedName := user.Name
		inputName := input.Name

		if inputName != expectedName {
			_, err := db.Query("INSERT INTO users (name, password, is_admin) VALUES ($1,$2,$3)", input.Name, hashedPassword, isAdmin)
			if err != nil {
				log.Println(err)
			}

		} else {
			log.Println("User", inputName, "allready exixts")
			return
		}
		w.Write([]byte(fmt.Sprintf("User %s created", input.Name)))

	}

}

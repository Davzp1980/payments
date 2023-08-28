package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//Поля в postman:
// "name"
// "password"

func CreateAdmin(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input Input

		json.NewDecoder(r.Body).Decode(&input)
		isAdmin := true
		hashedPassword, _ := HashedPassword(input.Password)

		_, err := db.Query("INSERT INTO users (name, password, is_admin) VALUES ($1,$2,$3)", input.Name, hashedPassword, isAdmin)
		if err != nil {
			log.Println(err)
			w.Write([]byte(fmt.Sprintf("User %s already exists", input.Name)))
			return
		}
		w.Write([]byte(fmt.Sprintf("User %s created", input.Name)))
	}

}

func CreateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input Input
		//var user User

		json.NewDecoder(r.Body).Decode(&input)
		isAdmin := false
		hashedPassword, _ := HashedPassword(input.Password)

		_, err := db.Query("INSERT INTO users (name, password, is_admin) VALUES ($1,$2,$3)", input.Name, hashedPassword, isAdmin)
		if err != nil {
			log.Println(err)
			w.Write([]byte(fmt.Sprintf("User %s already exists", input.Name)))
			return
		}

		/*if inputName != expectedName {
			_, err := db.Query("INSERT INTO users (name, password, is_admin) VALUES ($1,$2,$3)", input.Name, hashedPassword, isAdmin)
			if err != nil {
				log.Println(err)
			}

		} else {
			log.Println("User", inputName, "allready exixts")
			return
		}
		*/
		w.Write([]byte(fmt.Sprintf("User %s created", input.Name)))

	}

}

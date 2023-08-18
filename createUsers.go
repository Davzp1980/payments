package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

func CreateAdmin(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var admin User

		json.NewDecoder(r.Body).Decode(&admin)
		isAdmin := true
		hashedPassword, _ := HashedPassword(admin.PasswordHash)

		err := db.QueryRow("INSERT INTO users (name, password, is_admin) VALUES ($1,$2,$3) RETURNING id", admin.Name, hashedPassword, isAdmin).Scan(&admin.ID)
		if err != nil {
			log.Println(err)
		}
		json.NewEncoder(w).Encode(admin.PasswordHash)
	}

}

func CreateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input Input
		var user User

		json.NewDecoder(r.Body).Decode(&input)

		hashedPassword, _ := HashedPassword(user.PasswordHash)
		isAdmin := false

		err := db.QueryRow("SELECT * FROM users WHERE name=$1", input.Name).Scan(
			&user.ID, &user.Name, &user.PasswordHash, &user.IsAdmin)
		if err != nil {
			log.Println(err)
		}
		expectedName := user.Name
		inputName := input.Name

		if inputName != expectedName {
			err = db.QueryRow("INSERT INTO users (name, password,is_admin) VALUES ($1,$2,$3) RETURNING id", input.Name, hashedPassword, isAdmin).Scan(&user.ID)
			if err != nil {
				log.Println(err)
			}
			json.NewEncoder(w).Encode(input.Password)
		} else {
			log.Println("User", inputName, "allready exixts")
		}

	}

}

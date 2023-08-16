package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

func CreateAdmin(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var admin Admin

		json.NewDecoder(r.Body).Decode(&admin)

		hashedPassword, _ := HashedPassword(admin.PasswordHash)

		err := db.QueryRow("INSERT INTO admin (name, passwordHash) VALUES ($1,$2) RETURNING id", admin.Name, hashedPassword).Scan(&admin.ID)
		if err != nil {
			log.Println(err)
		}
		json.NewEncoder(w).Encode(admin)
	}

}

func CreateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var user User

		json.NewDecoder(r.Body).Decode(&user)

		hashedPassword, _ := HashedPassword(user.PasswordHash)

		err := db.QueryRow("SELECT * FROM users WHERE name=$1", user.Name).Scan(&user.ID, &user.Name, &user.PasswordHash)
		if err != nil {
			err = db.QueryRow("INSERT INTO admin (name, passwordHash) VALUES ($1,$2) RETURNING id", user.Name, hashedPassword).Scan(&user.ID)
			if err != nil {
				log.Println(err)
			}
			json.NewEncoder(w).Encode(user)
		} else {
			log.Println("Users allready exists")
		}

	}

}

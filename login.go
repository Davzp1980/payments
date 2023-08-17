package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("My_key")

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func Login(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input Input
		var user User

		json.NewDecoder(r.Body).Decode(&input)

		err := db.QueryRow("SELECT name, password FROM users WHERE name=$1", input.Name).Scan(
			&user.Name, &user.PasswordHash)
		if err != nil {
			log.Println(err)
		}

		if !CheckPassword(input.PasswordHash, user.PasswordHash) || input.Name != user.Name {
			http.Redirect(w, r, "/login", http.StatusUnauthorized)
			log.Println("Wrong password or user name")
			return
		}
		expirationTime := time.Now().Add(5 * time.Minute)

		claims := &Claims{
			Username: input.Name,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(expirationTime),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return

		}

		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})
		w.Write([]byte(fmt.Sprintf("Welcome %s", input.Name)))

	}
}

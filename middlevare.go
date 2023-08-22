package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type Middlevare interface {
	GetName(name string) bool
}

type PostgresRepo struct {
	db *sql.DB
}

var admin User

func HashedPassword(password string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func LogginingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/login" && r.URL.Path != "/createadmin" && r.URL.Path != "/createuser" {

			c, err := r.Cookie("token")
			if err != nil {
				if err == http.ErrNoCookie {
					w.WriteHeader(http.StatusUnauthorized)
					return
				}
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			tokenString := c.Value

			claims := &Claims{}

			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				return jwtKey, nil
			})

			if err != nil {
				if err == jwt.ErrSignatureInvalid {
					w.WriteHeader(http.StatusUnauthorized)
					return
				}
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			if !token.Valid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

func (g *PostgresRepo) AdminLogginingMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/login" && r.URL.Path != "/createadmin" {
			c, err := r.Cookie("token")
			if err != nil {
				if err == http.ErrNoCookie {
					w.WriteHeader(http.StatusUnauthorized)
					return
				}
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			tokenString := c.Value

			claims := &Claims{}

			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				return jwtKey, nil
			})
			if err != nil {
				if err == jwt.ErrSignatureInvalid {
					w.WriteHeader(http.StatusUnauthorized)
					return
				}
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			if g.GetName(claims.Username) == true {
				w.Write([]byte(fmt.Sprintf("User authorizated %s", claims.Username)))
			} else {
				log.Println("Not enough use rights lose ")
				w.WriteHeader(http.StatusUnauthorized)
			}
			if !token.Valid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

func (g *PostgresRepo) GetName(name string) bool {

	err := g.db.QueryRow("SELECT is_admin FROM users WHERE name=$1", name).Scan(&admin.IsAdmin)
	if err != nil {
		log.Println("User does not exists")
		return false
	}

	if !admin.IsAdmin {
		log.Println("Not enough use rights lose ")

		return false
	} else {
		log.Println("Authorizated")

		return true
	}
}

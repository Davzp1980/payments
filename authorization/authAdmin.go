package authorization

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

func CreateAdmin(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var Admin InputAdmin

		json.NewDecoder(r.Body).Decode(&Admin)

		hashedPassword, _ := HashedPassword(Admin.Password)

		err := db.QueryRow("INSERT INTO admin (name, passwordHash) VALUES ($1,$2) RETURNING id", Admin.Name, hashedPassword)
		if err != nil {
			log.Println(err)
		}
		json.NewEncoder(w).Encode(Admin)
	}

}

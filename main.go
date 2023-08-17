package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("BD is not opend")
	}
	defer db.Close()

	CreateNewDB(db)

	router := mux.NewRouter()

	router.HandleFunc("/createadmin", CreateAdmin(db)).Methods("POST")
	router.HandleFunc("/login", Login(db)).Methods("POST")
	router.HandleFunc("/createuser", CreateUser(db)).Methods("POST")

	router.Use(LogginingMiddleware)

	log.Fatal(http.ListenAndServe(":8000", router))

}

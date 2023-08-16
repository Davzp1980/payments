package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"payments/authorization"

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

	router.HandleFunc("/createAdmin", authorization.CreateAdmin(db))

	log.Fatal(http.ListenAndServe(":8000", router))

}

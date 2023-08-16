package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("BD did not opend")
	}
	defer db.Close()

	CreateNewDB_Admins_Users(db)

	router := mux.NewRouter()

	log.Fatal(http.ListenAndServe(":8000", router))

}

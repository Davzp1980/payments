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
		log.Println(err)
	}
	defer db.Close()

	CreateNewDB(db)

	router := mux.NewRouter()

	log.Fatal(http.ListenAndServe(":8000", router))

}

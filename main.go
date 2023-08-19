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

	//router.Use(AdminLogginingMiddleware)
	router.HandleFunc("/blockuser", BlockUser(db)).Methods("POST")
	router.HandleFunc("/unblockuser", UnBlockUser(db)).Methods("POST")
	router.HandleFunc("/createadmin", CreateAdmin(db)).Methods("POST")
	router.HandleFunc("/login", Login(db)).Methods("POST")
	router.HandleFunc("/createuser", CreateUser(db)).Methods("POST")
	router.HandleFunc("/changepassword", ChangeUserPassword(db)).Methods("POST")

	router.Use(LogginingMiddleware)
	router.HandleFunc("/createaccount", CreateAccount(db)).Methods("POST")
	router.HandleFunc("/createapayment", CreatePayment(db)).Methods("POST")

	router.HandleFunc("/blockaccount", BlockAccount(db)).Methods("POST")
	router.HandleFunc("/unblockaccount", UnBlockAccount(db)).Methods("POST")

	router.HandleFunc("/getaccountid", GetAccountsById(db)).Methods("GET")
	router.HandleFunc("/getaccountiban", GetAccountsByIban(db)).Methods("GET")
	router.HandleFunc("/getaccountbalance", GetAccountsByBalance(db)).Methods("GET")

	router.HandleFunc("/getpaymentid", GetPaymentsById(db)).Methods("GET")
	router.HandleFunc("/getpaymentdate", GetPaymentsDate(db)).Methods("GET")

	router.HandleFunc("/logout", Logout).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))

}

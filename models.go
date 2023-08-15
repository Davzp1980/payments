package main

import "time"

type Admin struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	PasswordHash string `json:"passwordhash"`
}

type User struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	PasswordHash string `json:"passwordhash"`
}

type Payment struct {
	Number        string    `json:"number"`
	AmountPayment int       `json:"amountPayment"`
	Date          time.Time `json:"date"`
}

type UserAccount struct {
	ID             int    `json:"id"`
	Iban           string `json:"iban"`
	AccountBalance int    `json:"accountBalance"`
}

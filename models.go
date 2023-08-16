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

type UserCabinet struct {
	ID             int    `json:"id"`
	UserName       string `json:"userName"`
	Iban           string `json:"iban"`
	AccountBalance int    `json:"accountBalance"`
}

type Payment struct {
	ID            int       `json:"id"`
	UserName      string    `json:"userName"`
	NumberPayment string    `json:"numberPayment"`
	AmountPayment int       `json:"amountPayment"`
	Date          time.Time `json:"date"`
}

type InputAdmin struct {
	Name     string `json:"name"`
	Password string `json:"passwordhash"`
}

type InputUser struct {
	Name     string `json:"name"`
	Password string `json:"passwordhash"`
}

type InputPayment struct {
	UserName      string `json:"userName"`
	AmountPayment int    `json:"amountPayment"`
}

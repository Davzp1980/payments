package main

import "time"

type User struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	PasswordHash string `json:"password_hash"`
	IsAdmin      bool   `json:"is_admin"`
}

type Account struct {
	ID      int    `json:"id"`
	UserId  int    `json:"user_id"`
	Iban    string `json:"iban"`
	Balance int    `json:"balance"`
}

type Payment struct {
	ID            int       `json:"id"`
	AccountId     string    `json:"account_id"`
	AmountPayment int       `json:"amount_payment"`
	Date          time.Time `json:"date"`
	Reciever      string    `json:"reciever"`
}

type Input struct {
	Name         string `json:"name"`
	PasswordHash string `json:"password_hash"`
}

type InputPayment struct {
	PayerName     string `json:"payer_name"`
	ReceiverName  string `json:"receiver_name"`
	AmountPayment int    `json:"amount_payment"`
}

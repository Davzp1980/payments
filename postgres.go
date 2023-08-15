package main

import "database/sql"

func CreateNewDB_Admins_Users(db *sql.DB) error {

	CreateTablesQuery := `
		CREATE TABLE IF NOT EXISTS admin (
			id serial PRYMARY KEY,
			name VARCHAR (25) NOT NULL UNIQUE,
			passwordHash VARCHAR NOT NULL UNIQUE
		);
		CREATE TABLE IF NOT EXISTS user (
			id serial PRYMARY KEY,
			name VARCHAR (25) NOT NULL UNIQUE,
			password_hash VARCHAR NOT NULL UNIQUE
		);
		CREATE TABLE IF NOT EXISTS usercabinet (
			id serial PRYMARY KEY,
			user_name VARCHAR (25),
			iban VARCHAR,
			AccountBalance serial
		);
		CREATE TABLE IF NOT EXISTS payment (
			id serial PRYMARY KEY,
			user_name VARCHAR (25),
			number_payment VARCHAR,
			amount_payment serial,
			date TIMESTAMP WITH TIME ZONE
		);

	`

	_, err := db.Exec(CreateTablesQuery)
	return err
}

func CreateNewDB_UserCab_Payment(db *sql.DB) error {

	CreateTablesQuery := `
		CREATE TABLE IF NOT EXISTS usercabinet (
			id serial PRYMARY KEY,
			user_name VARCHAR (25),
			iban VARCHAR,
			AccountBalance serial
		);
		CREATE TABLE IF NOT EXISTS payment (
			id serial PRYMARY KEY,
			user_name VARCHAR (25),
			number_payment VARCHAR,
			amount_payment serial,
			date TIMESTAMP WITH TIME ZONE
		);

	`

	_, err := db.Exec(CreateTablesQuery)
	return err
}

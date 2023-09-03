package main

import (
	"bytes"
	"database/sql"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/magiconair/properties/assert"
)

func Test_CreateAdmin(t *testing.T) {
	//docker run --name=test -e POSTGRES_PASSWORD='qwerty' -p 5432:5432 -d --rm postgres
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres password=qwerty dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	CreateTablesQuery := `
		CREATE TABLE IF NOT EXISTS users (
			id serial PRIMARY KEY,
			name VARCHAR (25) NOT NULL UNIQUE,
			password VARCHAR NOT NULL UNIQUE,
			is_admin boolean,
			blocked boolean DEFAULT false
		);
	`
	_, err = db.Exec(CreateTablesQuery)

	tests := []struct {
		name               string
		inputBody          string
		inputUser          User
		expectedStatusCode int
		//expectedResponseBody string
	}{
		{
			name:      "OK",
			inputBody: `{"name":"Test", "password":"qwerty"}`,
			inputUser: User{
				Name:         "Test",
				PasswordHash: "qwerty",
			},
			expectedStatusCode: 200,
			//expectedResponseBody: `{"User Test created"}`,
		},
		{
			name:      "fail",
			inputBody: `{"name":"Test", "password":"qwerty"}`,
			inputUser: User{
				Name:         "Test",
				PasswordHash: "qwerty",
			},
			expectedStatusCode: 400,
			//expectedResponseBody: `{"User Test already exists"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			r := mux.NewRouter()
			r.HandleFunc("/createadmin", CreateAdmin(db))

			// Создаем тестовый HTTP ResponseWriter и запускаем обработчик
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/createadmin", bytes.NewBufferString(test.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			//assert.Equal(t, w.Code, test.expectedResponseBody)
		})
	}

}

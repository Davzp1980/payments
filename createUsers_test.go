package main

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_CreateUser(t *testing.T) {

	request := http.Request{
		Method: "POST",
		DB:     sql.DB,
	}

	rw := httptest.NewRecorder()

	CreateUser(db * sql.DB)

}

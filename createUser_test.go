package main

import (
	"bytes"
	"database/sql"
	"net/http/httptest"
	"testing"

	"github.com/a1div0/fakedb"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/magiconair/properties/assert"
)

func Test_CreateAdmin(t *testing.T) {
	fc := &fakedb.FakeConnector{}
	db := sql.OpenDB(fc)
	if db.Driver() != fakedb.Fdriver {
		t.Error("OpenDB should return the driver of the Connector")
		return
	}
	defer db.Close()

	tests := []struct {
		name                 string
		inputBody            string
		inputUser            User
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "OK",
			inputBody: `{"name":"Test", "password":"qwerty"}`,
			inputUser: User{
				Name:         "Test",
				PasswordHash: "qwerty",
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"User Test created"}`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			r := mux.NewRouter()
			r.HandleFunc("/createadmin", CreateAdmin(db))

			// Создаем фейковый HTTP ResponseWriter и запускаем обработчик
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/createadmin", bytes.NewBufferString(test.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Code, test.expectedResponseBody)
		})
	}

}

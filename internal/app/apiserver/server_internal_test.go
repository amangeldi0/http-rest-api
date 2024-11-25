package apiserver

import (
	"bytes"
	"encoding/json"
	"github.com/amangeldi0/http-rest-api/internal/app/model"
	"github.com/amangeldi0/http-rest-api/internal/app/store/teststore"
	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	sessionsStore = sessions.NewCookieStore([]byte("something-very-secret"))
)

func TestServer_HandleUsersCreate(t *testing.T) {

	s := newServer(teststore.New(), sessionsStore)

	testsCases := []struct {
		name         string
		expectedCode int
		payload      interface{}
	}{
		{
			name:         "valid",
			expectedCode: http.StatusCreated,
			payload: map[string]string{
				"email":    "user@example.org",
				"password": "password",
			},
		},

		{
			name:         "invalid payload",
			expectedCode: http.StatusBadRequest,
			payload:      "invalid",
		},

		{
			name:         "invalid params",
			expectedCode: http.StatusUnprocessableEntity,
			payload: map[string]string{
				"email":    "invalid",
				"password": "",
			},
		},
	}

	for _, testCase := range testsCases {
		t.Run(testCase.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(testCase.payload)
			req, _ := http.NewRequest("POST", "/users", b)

			s.ServeHTTP(rec, req)

			assert.Equal(t, testCase.expectedCode, rec.Code)

		})
	}

}

func TestServer_HandleSessionsCreate(t *testing.T) {
	u := model.TestUser(t)
	s := teststore.New()
	srv := newServer(s, sessionsStore)

	s.User().Create(u)

	testsCases := []struct {
		name         string
		expectedCode int
		payload      interface{}
	}{
		{
			name:         "valid",
			expectedCode: http.StatusOK,
			payload: map[string]string{
				"email":    u.Email,
				"password": u.Password,
			},
		},

		{
			name:         "invalid email",
			expectedCode: http.StatusUnauthorized,
			payload: map[string]string{
				"email":    "aa",
				"password": u.Password,
			},
		},

		{
			name:         "invalid password",
			expectedCode: http.StatusUnauthorized,
			payload: map[string]string{
				"email":    u.Email,
				"password": "gg",
			},
		},
		{
			name:         "valid",
			expectedCode: http.StatusBadRequest,
			payload:      "invalid",
		},
	}

	for _, testCase := range testsCases {
		t.Run(testCase.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(testCase.payload)
			req, _ := http.NewRequest("POST", "/sessions", b)

			srv.ServeHTTP(rec, req)

			assert.Equal(t, testCase.expectedCode, rec.Code)

		})
	}
}

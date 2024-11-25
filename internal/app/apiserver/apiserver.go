package apiserver

import (
	"database/sql"
	"github.com/amangeldi0/http-rest-api/internal/app/store/sqlstore"
	"github.com/gorilla/sessions"
	"net/http"
)

func Start(conf *Config) error {
	db, err := newDB(conf.DatabaseURL)

	if err != nil {
		return err
	}

	defer db.Close()
	store := sqlstore.New(db)
	sessionsStore := sessions.NewCookieStore([]byte(conf.SessionKey))
	srv := newServer(store, sessionsStore)

	return http.ListenAndServe(conf.BindAddr, srv)
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

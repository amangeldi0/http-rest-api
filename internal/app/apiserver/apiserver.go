package apiserver

import (
	"database/sql"
	"github.com/amangeldi0/http-rest-api/internal/app/store/sqlstore"
	"net/http"
)

func Start(conf *Config) error {
	db, err := newDB(conf.DatabaseURL)

	if err != nil {
		return err
	}

	defer db.Close()
	store := sqlstore.New(db)
	srv := newServer(store)

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

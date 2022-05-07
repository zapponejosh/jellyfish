package dbops

import (
	"context"
	"database/sql"

	"github.com/zapponejosh/jellyfish/internal/settings"
)

type DB struct {
	db *sql.DB
}

type Test struct {
	testStr string
}

func (d DB) Close() error {
	return nil
}

func New(s *settings.Settings) (*DB, error) {
	db, err := sql.Open("postgres", s.DBSettings.DSN())
	if err != nil {
		return nil, err
	}
	return &DB{db: db}, nil
}

func (d DB) GetTest(ctx context.Context) (*Test, error) {
	rows, err := d.db.Query("select * from test")
	if err != nil {
		panic(err.Error())
	}
	var testStr string
	for rows.Next() {
		err = rows.Scan(&testStr)
		if err != nil {
			return nil, err
		}
	}
	return &testStr, nil
}

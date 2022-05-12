package dbops

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/zapponejosh/jellyfish/internal/models"
	"github.com/zapponejosh/jellyfish/internal/settings"
)

type DB struct {
	db *sql.DB
}

func (d DB) Close() error {
	return nil
}

func New(s *settings.Settings) (*DB, error) {
	fmt.Println(s.DBSettings.DSN())
	db, err := sql.Open("postgres", s.DBSettings.DSN())
	if err != nil {
		return nil, err
	}
	return &DB{db: db}, nil
}

func (d DB) GetTest(ctx context.Context) (*models.Test, error) {
	rows, err := d.db.Query("select * from test")
	if err != nil {
		panic(err.Error())
	}
	var testR models.Test

	for rows.Next() {
		err = rows.Scan(&testR.TestStr)
		if err != nil {
			return nil, err
		}
	}
	return &testR, nil
}

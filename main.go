package main

import (
	"fmt"

	"os"

	"github.com/zapponejosh/jellyfish/internal/dbops"
	"github.com/zapponejosh/jellyfish/internal/server"
	"github.com/zapponejosh/jellyfish/internal/settings"
)

func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}

func run() error {
	appSettings, err := settings.New()
	if err != nil {
		fmt.Println("error getting settings", err)
		return err
	}
	fmt.Println(appSettings)

	db, err := dbops.New(appSettings)
	if err != nil {
		fmt.Println("error connecting to db", err)
		return err
	}
	defer db.Close()

	svr, err := server.New(appSettings, db)
	if err != nil {
		fmt.Println("error starting server", err)
		return err
	}

	if err := svr.ListenAndServe(); err != nil {
		fmt.Println("error running server", err)
		return err
	}
	return nil
}

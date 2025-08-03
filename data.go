package main

import (
	"fyne.io/fyne/v2"

	"database/sql"
	"log"
	"os"
	"path"
)

const (
	SQLMigrations string = `
	
	`
)

func loadDatabases(app fyne.App) {
	newDb := false

	if _, err := os.Stat(path.Join(app.Storage().RootURI().Path(), "app.db")); os.IsNotExist(err) {
		// create database
		newDb = true
	}

	db, err := sql.Open("sqlite3", path.Join(app.Storage().RootURI().Path(), "app.db"))

	if err != nil {
		log.Println("Error opening database:", err)
		return
	}

	appState.db = db

	if newDb {
		_, err = appState.db.Exec(SQLMigrations)

		if err != nil {
			log.Println("Error setting up database:", err)
			return
		}
	}
}

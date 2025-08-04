package main

import (
	"fyne.io/fyne/v2"

	"database/sql"
	"log"
	"os"
	"path"

	_ "github.com/mattn/go-sqlite3"
)

const (
	SQLMigrations string = `
	create table members (
		id integer not null primary key,
		name text unique not null,
		pronouns text,
		avatar_path text,
		color text,
		proxy_tags text
	);
	
	create table messages (
		id integer not null primary key autoincrement,
		member_id integer not null,
		message text not null,
		timestamp text not null,
		create_at datetime default current_timestamp,
		foreign key (member_id) references members (id)
	);
	`
)

type ProxyTag struct {
	Prefix *string `json:"prefix,omitempty"`
	Suffix *string `json:"suffix,omitempty"`
}

type Proxies []ProxyTag

func loadDatabases(app fyne.App, state *AppState) {
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

	(*state).db = db

	if newDb {
		_, err = state.db.Exec(SQLMigrations)

		if err != nil {
			log.Println("Error setting up database:", err)
			return
		}
	}
}

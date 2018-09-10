package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStmt := `
  create table foo (id integer not null primary key, name text);
	insert into foo (name) values ("test")
  `
	_, err = db.Exec(sqlStmt)
	if err != nil {
		panic(err)
	}
}

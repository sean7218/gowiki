package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "fmt"
)

func setupDatabase() {
	db, err := sql.Open("mysql", "dbname")
	
}
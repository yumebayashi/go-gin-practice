package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)


func openConnection() (*sql.DB, error) {
	return sql.Open("mysql", "root:rootpasswd@(localhost:3306)/adlive")
}


func getLiveUrl(db *sql.DB, key string) string{

	var url string

	query := "SELECT url FROM channel WHERE id = ?"

	rows, qerr := db.Query(query, key)

	if qerr != nil {
		log.Fatal("query error: %v", qerr)
	}

	for rows.Next() {
		var c1 string
		if berr := rows.Scan(&c1); berr != nil {
			log.Fatal("scan erro: %v", berr)
		}
		url = c1
	}
	rows.Close()
	return url
}

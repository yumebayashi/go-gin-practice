package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)


func openConnection() (*sql.DB, error) {
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASS")
	url := os.Getenv("MYSQL_URL")

	return sql.Open("mysql", user + ":" + pass + "@(" + url + ":3306)/adlive")
}


func getLiveUrl(db *sql.DB, key string) string {

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

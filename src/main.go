package main

import (
	"github.com/gin-gonic/gin"
//	"database/sql"
	_ "github.com/go-sql-driver/mysql"
//	"fmt"
	"log"
	"net/http"
)

func main() {

//	db, err := sql.Open("mysql", "root:rootpasswd@/go-test")
//	if err != nil {
//		log.Fatal("open erro: %v", err)
//	}
//	defer db.Close()
//
//	rows, qerr := db.Query("SELECT * FROM `live`")
//
//	defer rows.Close()
//	if qerr != nil {
//		log.Fatal("query error: %v", qerr)
//	}
//
//	for rows.Next() {
//		var c1 int
//		var c2 string
//		var c3 string
//		if berr := rows.Scan(&c1, &c2, &c3); berr != nil {
//			log.Fatal("scan erro: %v", berr)
//		}
//		fmt.Println(c1, c2, c3)
//	}

	db, err := openConnection()

	if err != nil {
		log.Fatalf("データベースの接続に失敗しました。: %v", err)
	}

	defer db.Close()

	r := gin.Default()
	r.GET("/live/:liveId", func(c *gin.Context) {
		liveId := c.Param("liveId")
		c.JSON(http.StatusOK, gin.H{"url": getLiveUrl(db, liveId)})
	})

	// Listen and server on 0.0.0.0:8080
	r.Run(":8080")
}

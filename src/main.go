package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func main() {

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

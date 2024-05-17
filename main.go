package main

import (
	"fmt"
	"github.com/AbassAdeyemi/bookmarks/internal/config"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	cfg, _ := config.GetConfig("config.json")
	r := gin.Default()
	r.GET("/hello", hello)
	log.Fatal(r.Run(fmt.Sprintf(":%d", cfg.Db.Port)))
}

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}

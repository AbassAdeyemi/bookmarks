package main

import (
	"fmt"
	"github.com/AbassAdeyemi/bookmarks/internal/config"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.GetConfig("config.json")
	if err != nil {
		log.Fatal(err)
	}
	logger := config.NewLogger(cfg)
	logger.Infof("Application is running on %d", cfg.ServerPort)
	r := gin.Default()
	r.GET("/hello", hello)
	log.Fatal(r.Run(fmt.Sprintf(":%d", cfg.ServerPort)))
}

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}

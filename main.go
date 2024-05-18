package main

import (
	"github.com/AbassAdeyemi/bookmarks/cmd"
	"github.com/AbassAdeyemi/bookmarks/internal/config"
	"log"
)

func main() {
	cfg, err := config.GetConfig("config.json")
	if err != nil {
		log.Fatal(err)
	}
	app := cmd.NewApp(cfg)
	app.Run()
}

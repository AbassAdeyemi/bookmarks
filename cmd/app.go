package cmd

import (
	"fmt"
	"github.com/AbassAdeyemi/bookmarks/internal/api"
	"github.com/AbassAdeyemi/bookmarks/internal/config"
	"github.com/AbassAdeyemi/bookmarks/internal/domain"
	"github.com/gin-gonic/gin"
	"log"
)

type App struct {
	Router *gin.Engine
	cfg    config.AppConfig
}

func NewApp(cfg config.AppConfig) *App {
	logger := config.NewLogger(cfg)
	db := config.GetDb(cfg, logger)

	repo := domain.NewBookmarkRepository(db, logger)
	handler := api.NewBookmarkController(repo, logger)

	router := gin.Default()
	router.GET("/api/bookmarks", handler.GetAll)
	router.POST("/api/bookmarks", handler.Create)
	router.GET("/api/bookmarks/:id", handler.GetByID)
	router.PUT("/api/bookmarks/:id", handler.Update)
	router.DELETE("/api/bookmarks/:id", handler.Delete)

	return &App{
		Router: router,
		cfg:    cfg,
	}
}

func (app App) Run() {
	log.Fatalln(app.Router.Run(fmt.Sprintf(":%d", app.cfg.ServerPort)))
}

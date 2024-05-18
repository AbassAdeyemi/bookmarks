package api

import (
	"github.com/AbassAdeyemi/bookmarks/internal/config"
	"github.com/AbassAdeyemi/bookmarks/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BookmarkController struct {
	repo   domain.BookmarkRepository
	logger *config.Logger
}

func NewBookmarkController(repo domain.BookmarkRepository, logger *config.Logger) BookmarkController {
	return BookmarkController{
		repo:   repo,
		logger: logger,
	}
}

func (b BookmarkController) GetAll(c *gin.Context) {
	b.logger.Infof("Getting all bookmarks")
	ctx := c.Request.Context()
	bookmarks, err := b.repo.GetAll(ctx)
	if err != nil {
		if err != nil {
			b.logger.Errorf("Error: %v", err)
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to fetch bookmarks",
		})
		return
	}
	c.JSON(http.StatusOK, bookmarks)
}

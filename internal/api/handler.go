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

type CreateBookmarkRequest struct {
	Title string `json:"title" binding:"required"`
	Url   string `json:"url" binding:"required,url"`
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

func (b BookmarkController) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var model CreateBookmarkRequest
	if err := c.ShouldBindJSON(&model); err != nil {
		b.respondWithError(c, http.StatusBadRequest, err, "Invalid request payload")
		return
	}
	b.logger.Infof("Creating bookmark for URL, %v", model.Url)
	bookmark := domain.Bookmark{
		Title: model.Title,
		Url:   model.Url,
	}

	savedBookmark, err := b.repo.Create(ctx, bookmark)
	if err != nil {
		b.respondWithError(c, http.StatusInternalServerError, err, "Failed to save bookmark")
	}

	c.JSON(http.StatusOK, savedBookmark)
}

func (b BookmarkController) respondWithError(c *gin.Context, code int, err error, errMsg string) {
	if err != nil {
		b.logger.Errorf("Error: %v", err)
	}
	c.AbortWithStatusJSON(code, gin.H{
		"error": errMsg,
	})
}

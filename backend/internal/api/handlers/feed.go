package handlers

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"drakemaye/backend/internal/storage"

	"github.com/gin-gonic/gin"
)

type FeedHandler struct{ db *sql.DB }

func NewFeedHandler(db *sql.DB) FeedHandler { return FeedHandler{db: db} }

func (h FeedHandler) List(c *gin.Context) {
	posts, err := storage.ListFeedPosts(h.db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, posts)
}

func (h FeedHandler) Like(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	post, err := storage.LikeFeedPost(h.db, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, post)
}

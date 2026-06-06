package handlers

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"drakemaye/backend/internal/models"
	"drakemaye/backend/internal/services"
	"drakemaye/backend/internal/storage"

	"github.com/gin-gonic/gin"
)

type LogsHandler struct{ db *sql.DB }

func NewLogsHandler(db *sql.DB) LogsHandler { return LogsHandler{db: db} }

func (h LogsHandler) List(c *gin.Context) {
	logs, err := storage.ListLogs(h.db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, logs)
}

func (h LogsHandler) Get(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	log, err := storage.GetLog(h.db, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "log not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, log)
}

func (h LogsHandler) Create(c *gin.Context) {
	var log models.DrinkLog
	if err := c.ShouldBindJSON(&log); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	if log.UserID == 0 {
		// TODO: Replace mock user fallback with authenticated user context.
		log.UserID = 1
	}
	if err := services.ValidateLogInput(log); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	services.EnrichLogStandardDrinks(&log)

	username, err := storage.GetUserName(h.db, log.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to resolve user"})
		return
	}

	created, err := storage.CreateLogAndFeedPost(h.db, log, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, created)
}

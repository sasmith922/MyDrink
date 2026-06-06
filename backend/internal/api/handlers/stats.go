package handlers

import (
	"database/sql"
	"net/http"
	"time"

	"drakemaye/backend/internal/services"
	"drakemaye/backend/internal/storage"

	"github.com/gin-gonic/gin"
)

type StatsHandler struct{ db *sql.DB }

func NewStatsHandler(db *sql.DB) StatsHandler { return StatsHandler{db: db} }

func (h StatsHandler) AllTime(c *gin.Context) {
	logs, err := storage.ListLogs(h.db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, services.BuildStatsSummary("all-time", logs))
}

func (h StatsHandler) Weekly(c *gin.Context) {
	since := time.Now().AddDate(0, 0, -7)
	logs, err := storage.ListLogsSince(h.db, since)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, services.BuildStatsSummary("weekly", logs))
}

func (h StatsHandler) Monthly(c *gin.Context) {
	since := time.Now().AddDate(0, -1, 0)
	logs, err := storage.ListLogsSince(h.db, since)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, services.BuildStatsSummary("monthly", logs))
}

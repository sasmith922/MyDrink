package api

import (
	"database/sql"
	"net/http"

	"drakemaye/backend/internal/api/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: false,
	}))

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	drinks := handlers.NewDrinksHandler(db)
	logs := handlers.NewLogsHandler(db)
	feed := handlers.NewFeedHandler(db)
	stats := handlers.NewStatsHandler(db)
	profile := handlers.NewProfileHandler(db)

	r.GET("/drinks", drinks.List)
	r.GET("/drinks/:id", drinks.Get)
	r.POST("/drinks", drinks.Create)

	r.GET("/logs", logs.List)
	r.POST("/logs", logs.Create)
	r.GET("/logs/:id", logs.Get)

	r.GET("/feed", feed.List)
	r.POST("/feed/:id/like", feed.Like)

	r.GET("/stats", stats.AllTime)
	r.GET("/stats/weekly", stats.Weekly)
	r.GET("/stats/monthly", stats.Monthly)

	r.GET("/profile", profile.Get)
	r.PUT("/profile/goals", profile.UpdateGoals)

	return r
}

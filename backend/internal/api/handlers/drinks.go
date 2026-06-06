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

type DrinksHandler struct{ db *sql.DB }

func NewDrinksHandler(db *sql.DB) DrinksHandler { return DrinksHandler{db: db} }

func (h DrinksHandler) List(c *gin.Context) {
	drinks, err := storage.ListDrinks(h.db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, drinks)
}

func (h DrinksHandler) Get(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	drink, err := storage.GetDrink(h.db, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "drink not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, drink)
}

func (h DrinksHandler) Create(c *gin.Context) {
	var drink models.Drink
	if err := c.ShouldBindJSON(&drink); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	if err := services.ValidateDrinkInput(drink); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	created, err := storage.CreateDrink(h.db, drink)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, created)
}

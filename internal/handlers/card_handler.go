package handlers

import (
	"net/http"
	"strconv"

	"github.com/DestWish/cards/internal/models"
	"github.com/DestWish/cards/internal/service"
	"github.com/gin-gonic/gin"
)

type CardHandler struct {
	service *service.CardService
}

func NewCardHandler(service *service.CardService) *CardHandler {
	return &CardHandler{service: service}
}

func (h *CardHandler) RegisterRoutes(r *gin.Engine) {
	cards := r.Group("/api/cards")
	{
		cards.POST("", h.Create)
		cards.GET("", h.GetAll)
		cards.PUT("", h.Update)
	}
}

func (h *CardHandler) Create(c *gin.Context) {
	var req models.CreateCardRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	card, err := h.service.CreateCard(req)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Card already exists"})
		return
	}
	c.JSON(http.StatusCreated, card)
}

func (h *CardHandler) GetAll(c *gin.Context) {
	 uid, ok := parseUserId(c)
	 if !ok {
		return
	 }

	 cards, err := h.service.GetAllCards(uid)
	 if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	 }
	 c.JSON(http.StatusOK, cards)
}

func (h *CardHandler) Update(c *gin.Context) {
	var req models.UpdateCardRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	card, err := h.service.UpdateCard(req)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Update failed"})
		return
	}
	c.JSON(http.StatusAccepted, card)
}



func parseUserId(c *gin.Context) (uint, bool) {
	id, err := strconv.ParseUint(c.Param("user_id"), 10, 32)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user_id"})
		return 0, false
	}
	return uint(id), true
}





package handler

import (
	"net/http"
	"strconv"
	"subscription-service/model"
	"subscription-service/service"

	"github.com/gin-gonic/gin"
)

type SubscriptionHandler struct {
	service *service.SubscriptionService
}

func NewSubscriptionHandler(s *service.SubscriptionService) *SubscriptionHandler {
	return &SubscriptionHandler{service: s}
}

func (h *SubscriptionHandler) RegisterRoutes(r *gin.Engine) {
	r.POST("/subscriptions", h.Create)
	r.GET("/subscriptions", h.GetAll)
	r.GET("/subscriptions/:id", h.GetByID)
	r.PUT("/subscriptions/:id", h.Update)
	r.DELETE("/subscriptions/:id", h.Delete)
	r.GET("/subscriptions/total", h.Total)
}

func (h *SubscriptionHandler) Create(c *gin.Context) {
	var sub model.Subscription
	if err := c.ShouldBindJSON(&sub); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.Create(&sub); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sub)
}

func (h *SubscriptionHandler) GetAll(c *gin.Context) {
	subs, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, subs)
}

func (h *SubscriptionHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	sub, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}
	c.JSON(http.StatusOK, sub)
}

func (h *SubscriptionHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var sub model.Subscription
	if err := c.ShouldBindJSON(&sub); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.Update(id, &sub); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "updated"})
}

func (h *SubscriptionHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}

func (h *SubscriptionHandler) Total(c *gin.Context) {
	service := c.Query("service")
	userId := c.Query("userId")
	total, err := h.service.Total(service, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"total_price": total})
}

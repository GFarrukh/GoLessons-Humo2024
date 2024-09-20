package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"pharmacy-system/internal/model"
	"strconv"
)

func (h *Handler) GetBorrows(c *gin.Context) {
	borrows, err := h.service.GetBorrows()
	if err != nil {
		log.Printf("GetBorrows - h.service.GetBorrows error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("GetBorrows - borrows: %v", borrows)
	c.JSON(http.StatusOK, gin.H{"data": borrows})
}

func (h *Handler) GetBorrowByID(c *gin.Context) {
	idStr := c.Param("ID")

	if idStr == "" {
		log.Printf("GetBorrowByID - id is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("GetBorrowByID - strconv.Atoi error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	borrow, err := h.service.GetBorrowByID(id)
	if err != nil {
		log.Printf("GetBorrowByID - h.service.GetBorrowByID error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("GetBorrowByID - borrow: %v", borrow)
	c.JSON(http.StatusOK, gin.H{"data": borrow})
}

func (h *Handler) GetBorrowByUser(c *gin.Context) {
	idStr := c.Param("ID")
	if idStr == "" {
		log.Printf("GetBorrowByUser - id is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("GetBorrowByUser - strconv.Atoi error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	borrows, err := h.service.GetBorrowsByUser(id)
	if err != nil {
		log.Printf("GetBorrowByUser - h.service.GetBorrowsByUser error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("GetBorrowByUser - borrows: %v", borrows)
	c.JSON(http.StatusOK, gin.H{"data": borrows})
}

func (h *Handler) GetBorrowByMedication(c *gin.Context) {
	idStr := c.Param("ID")
	if idStr == "" {
		log.Printf("GetBorrowByMedication - id is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("GetBorrowByMedication - strconv.Atoi error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	borrows, err := h.service.GetBorrowsByMedication(id)
	if err != nil {
		log.Printf("GetBorrowByMedication - h.service.GetBorrowsByMedication error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("GetBorrowByMedication - borrows: %v", borrows)
	c.JSON(http.StatusOK, gin.H{"data": borrows})
}

func (h *Handler) CreateBorrow(c *gin.Context) {
	var borrow model.Borrow

	if err := c.BindJSON(&borrow); err != nil {
		log.Printf("CreateBorrow - c.BindJSON error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("CreateBorrow - data after binding: %v", borrow)

	// Получение ID пользователя из контекста
	userID, ok := c.Get("UserID")
	if !ok {
		log.Printf("CreateBorrow - user_id is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}
	borrow.UserID = userID.(int)

	createdBorrow, err := h.service.CreateBorrow(&borrow)
	if err != nil {
		log.Printf("CreateBorrow - h.service.CreateBorrow error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("CreateBorrow - createdBorrow: %v", createdBorrow)
	c.JSON(http.StatusOK, gin.H{"data": createdBorrow})
}

func (h *Handler) ReturnMedication(c *gin.Context) {
	idStr := c.Param("ID")
	if idStr == "" {
		log.Printf("ReturnMedication - id is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("ReturnMedication - strconv.Atoi error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, ok := c.Get("UserID")
	if !ok {
		log.Printf("ReturnMedication - UserID is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "UserID is required"})
		return
	}

	if err := h.service.ReturnMedication(userID.(int), id); err != nil {
		log.Printf("ReturnMedication - h.service.ReturnMedication error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("ReturnMedication - medication with id %d returned by user %d", id, userID)
	c.JSON(http.StatusOK, gin.H{"message": "Medication returned"})
}

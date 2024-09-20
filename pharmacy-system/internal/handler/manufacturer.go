package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"pharmacy-system/internal/model"
	"strconv"
)

func (h *Handler) GetManufacturer(c *gin.Context) {
	manufacturer, err := h.service.GetManufacturer()
	if err != nil {
		log.Printf("GetManufacturer - h.service.ListManufacturer error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("GetManufacturer - manufacturer: %v", manufacturer)
	c.JSON(http.StatusOK, gin.H{"data": manufacturer})
}

func (h *Handler) GetManufacturerByID(c *gin.Context) {
	idStr := c.Param("ID")

	if idStr == "" {
		log.Printf("GetManufacturerByID - id is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("GetManufacturerByID - strconv.Atoi error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	manufacturer, err := h.service.GetManufacturerByID(id)
	if err != nil {
		log.Printf("GetManufacturerByID - h.service.GetManufacturerByID error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("GetManufacturerByID - author: %v", manufacturer)
	c.JSON(http.StatusOK, gin.H{"data": manufacturer})
}

func (h *Handler) AddManufacturer(c *gin.Context) {
	var manufacturer model.Manufacturer

	if err := c.BindJSON(&manufacturer); err != nil {
		log.Printf("AddManufacturer - c.BindJSON error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("AddManufacturer - data after binding: %v", manufacturer)

	createManufacturer, err := h.service.CreateManufacturer(&manufacturer)
	if err != nil {
		log.Printf("AddManufacturer - h.service.CreateManufacturer error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("AddManufacturer - created manufacturer: %v", createManufacturer)
	c.JSON(http.StatusOK, gin.H{"data": createManufacturer})
}

func (h *Handler) UpdateManufacturer(c *gin.Context) {
	var manufacturer model.Manufacturer

	if err := c.BindJSON(&manufacturer); err != nil {
		log.Printf("UpdateManufacturer - c.BindJSON error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("UpdateManufacturer - data after binding: %v", manufacturer)

	updateManufacturer, err := h.service.EditManufacurer(&manufacturer)
	if err != nil {
		log.Printf("UpdateManufacturer - h.service.EditManufacturer error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("UpdateManufacturer - updated manufacturer: %v", updateManufacturer)
	c.JSON(http.StatusOK, gin.H{"data": updateManufacturer})
}

func (h *Handler) DeleteManufacturer(c *gin.Context) {
	idStr := c.Param("ID")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("DeleteManufacturer - strconv.Atoi error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := h.service.DeleteManufacturer(id); err != nil {
		log.Printf("DeleteManufacturer - h.service.DeleteManufacturer error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("DeleteManufacturer - manufacturer with id %d deleted", id)
	c.JSON(http.StatusOK, gin.H{"message": "Manufacturer deleted"})
}

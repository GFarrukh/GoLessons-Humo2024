package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"pharmacy-system/internal/model"
	"strconv"
)

func (h *Handler) GetMedications(c *gin.Context) {
	medications, err := h.service.ListMedications()
	if err != nil {
		log.Printf("GetMedications - h.service.ListMedications error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("GetMedications - data: %v", medications)
	c.JSON(http.StatusOK, gin.H{"data": medications})
}

func (h *Handler) GetMedicationByID(c *gin.Context) {
	idStr := c.Param("ID")

	if idStr == "" {
		log.Printf("GetMedicationByID - id is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("GetMedicationByID - strconv.Atoi error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	medication, err := h.service.FindMedication(id)
	if err != nil {
		log.Printf("GetMedicationByID - h.service.FindMedication error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("GetMedicationByID - medication: %v", medication)
	c.JSON(http.StatusOK, gin.H{"data": medication})
}

func (h *Handler) GetMedicationsByManufacturer(c *gin.Context) {
	idStr := c.Param("ID")

	if idStr == "" {
		log.Printf("GetMedicationsByManufacturer - id is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("GetMedicationsByManufacturer - strconv.Atoi error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	medications, err := h.service.GetMedicationsByManufacturer(id)
	if err != nil {
		log.Printf("GetMedicationsByManufacturer - h.service.GetMedicationsByManufacturer error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("GetMedicationsByManufacturer - medications: %v", medications)
	c.JSON(http.StatusOK, gin.H{"data": medications})
}

func (h *Handler) AddMedication(c *gin.Context) {
	var medication model.Medications

	if err := c.BindJSON(&medication); err != nil {
		log.Printf("AddMedication - c.BindJSON error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("AddMedication - data after binding: %v", medication)

	createMedication, err := h.service.CreateMedication(&medication)
	if err != nil {
		log.Printf("AddMedication - h.service.CreateMedication error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("AddMedication - created medication: %v", createMedication)
	c.JSON(http.StatusCreated, gin.H{"data": createMedication})
}

func (h *Handler) UpdateMedication(c *gin.Context) {
	var medication model.Medications

	if err := c.BindJSON(&medication); err != nil {
		log.Printf("UpdateMedication - c.BindJSON error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("UpdateMediation - data after binding: %v", medication)

	updatedMedication, err := h.service.EditMedication(&medication)
	if err != nil {
		log.Printf("UpdateMedication - h.service.EditMedication error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("UpdateMedication - updated medication: %v", updatedMedication)
	c.JSON(http.StatusOK, gin.H{"data": updatedMedication})
}

func (h *Handler) DeleteMedication(c *gin.Context) {
	idStr := c.Param("ID")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("DeleteMedication - strconv.Atoi error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := h.service.RemoveMedication(id); err != nil {
		log.Printf("DeleteMedication - h.service.RemoveMedication error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("DeleteMedication - medication with id %d deleted", id)
	c.JSON(http.StatusOK, gin.H{"message": "Medication deleted"})
}

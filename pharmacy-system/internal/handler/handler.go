package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pharmacy-system/internal/handler/middleware"
	"pharmacy-system/internal/service"
)

type Handler struct {
	router  *gin.Engine
	service *service.Service
}

func NewHandler(router *gin.Engine, s *service.Service) *Handler {
	return &Handler{
		router:  router,
		service: s,
	}
}

func (h *Handler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "The service is up and running"})
}

func (h *Handler) InitRoutes() {
	h.router.Use(
		middleware.CORS(),
		middleware.Recovery(),
	)

	{
		h.router.GET("/health", h.HealthCheck)
		h.router.POST("/users/add", h.SignUp)
		h.router.POST("/users/login", h.SignIn)
	}

	v1 := h.router.Group("/v1")
	v1.Use(
		middleware.Authenticate(),
	)

	bookGroup := v1.Group("/medications")
	{
		bookGroup.GET("", h.GetMedications)
		bookGroup.GET("/{id}", h.GetMedicationByID)
		bookGroup.POST("/add", h.AddMedication)
		bookGroup.DELETE("/delete/{id}", h.DeleteMedication)
		bookGroup.PUT("/update", h.UpdateMedication)
		bookGroup.GET("/manufacturer/{id}", h.GetMedicationsByManufacturer)
	}

	userGroup := v1.Group("/users")
	{
		userGroup.GET("", h.GetUsers)
		userGroup.GET("/{id}", h.GetUserByID)
		userGroup.DELETE("/delete/{id}", h.DeleteUser)
		userGroup.PUT("/update", h.UpdateUser)
	}

	authorGroup := v1.Group("/manufacturers")
	{
		authorGroup.GET("", h.GetManufacturer)
		authorGroup.GET("/{id}", h.GetManufacturerByID)
		authorGroup.POST("/add", h.AddManufacturer)
		authorGroup.DELETE("/delete/{id}", h.DeleteManufacturer)
		authorGroup.PUT("/update", h.UpdateManufacturer)
	}

	reviewGroup := v1.Group("/reviews")
	{
		reviewGroup.GET("", h.GetReviews)
		reviewGroup.GET("/{id}", h.GetReviewByID)
		reviewGroup.GET("/user/{id}", h.GetReviewsByUser)
		reviewGroup.GET("/medication/{id}", h.GetReviewsByMedication)
		reviewGroup.GET("/filter", h.GetReviewsByFilter)
		reviewGroup.POST("/add", h.AddReview)
		reviewGroup.DELETE("/delete/{id}", h.DeleteReview)
		reviewGroup.PUT("/update", h.UpdateReview)
	}

	profileGroup := v1.Group("/profiles")
	{
		profileGroup.GET("", h.ListProfiles)
		profileGroup.GET("/{id}", h.GetProfileByID)
		profileGroup.POST("/add", h.CreateProfile)
		profileGroup.DELETE("/delete/{id}", h.DeleteProfile)
		profileGroup.PUT("/update", h.UpdateProfile)
	}

	borrowGroup := v1.Group("/borrows")
	{
		borrowGroup.GET("", h.GetBorrows)
		borrowGroup.GET("/{id}", h.GetBorrowByID)
		borrowGroup.GET("/user/{id}", h.GetBorrowByUser)
		borrowGroup.GET("/medication/{id}", h.GetBorrowByMedication)
		borrowGroup.POST("/add", h.CreateBorrow)
		borrowGroup.PUT("/return/{id}", h.ReturnMedication)
	}
}

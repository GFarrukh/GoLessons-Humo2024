package model

import (
	"fmt"
	"time"
)

type Medications struct {
	MedicationID   int        `json:"MedicationID,omitempty" gorm:"primaryKey"` //ID лекарств
	NameMedication string     `json:"NameMedication,omitempty"`                 //Название лекарств
	ManufacturerID int        `json:"ManufacturerID,omitempty"`                 //ID Инструкции
	CreatedAt      *time.Time `json:"created_at,omitempty"`
	UpdatedAt      *time.Time `json:"updated_at,omitempty"`
}

type Reviews struct {
	ReviewID     int        `json:"ReviewId" gorm:"primaryKey"`
	UserID       int        `json:"UserId"`
	MedicationID int        `json:"MedicationID"`
	ReviewText   string     `json:"ReviewText"`
	Rating       float64    `json:"Rating"`
	ReviewDate   *time.Time `json:"Review_date,omitempty"`
}

type ReviewFilter struct {
	ReviewID    int
	BookID      int
	UserID      int
	CountOnPage int
	Page        int
	DateFrom    *time.Time
	DateTo      *time.Time
}

// ValidateReviewFilter проверяет корректность значений фильтра
func (r *ReviewFilter) ValidateReviewFilter(filter ReviewFilter) error {
	if filter.CountOnPage < 0 {
		return fmt.Errorf("CountOnPage must be non-negative")
	}
	if filter.Page < 0 {
		return fmt.Errorf("page must be non-negative")
	}
	if filter.DateFrom != nil && filter.DateTo != nil && filter.DateFrom.After(*filter.DateTo) {
		return fmt.Errorf("DateFrom cannot be after DateTo")
	}
	return nil
}

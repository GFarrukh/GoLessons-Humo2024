package repository

import (
	"fmt"
	"log"
	"pharmacy-system/internal/model"
)

func (r *Repository) AddMedication(m *model.Medications) (*model.Medications, error) {
	// insert into Medication (Leaflet, ManufacturerID) values ('1 бори', 1)
	result := r.db.Create(&m)
	if result.Error != nil {
		log.Printf("AddMedication: Failed to add Medication: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to add Medication: %v\n", result.Error)
	}

	return m, nil
}

func (r *Repository) GetMedication() ([]model.Medications, error) {
	var m []model.Medications

	// select * from Medication;
	err := r.db.Find(&m).Error
	if err != nil {
		log.Printf("GetMedication: Failed to get Medications: %v\n", err)
		return nil, fmt.Errorf("cannon find Medications with error: %v", err)
	}

	return m, nil
}

func (r *Repository) GetMedicationByID(MedicationID int) (*model.Medications, error) {
	var m model.Medications

	// select * from Medication where MedicationID = MedicationID
	err := r.db.Where("MedicationID = ?", MedicationID).First(&m).Error
	if err != nil {
		log.Printf("GetMedicationByID: Failed to get Medication: %v\n", err)
		return nil, fmt.Errorf("cannot find Medication with error: %v", err)
	}

	return &m, nil
}

func (r *Repository) GetMedicationByManufacturer(ManufacturerID int) ([]model.Medications, error) {
	var m []model.Medications

	// select * from Medication where ManufacturerID = ManufacturerID
	err := r.db.Where("ManufacturerID = ?", ManufacturerID).Find(&m).Error
	if err != nil {
		log.Printf("GetMedicationByAuthor: Failed to get Medication: %v\n", err)
		return nil, fmt.Errorf("Cannot find Medication by ManufacturerID with error: %v", err)
	}

	return m, nil
}

func (r *Repository) GetMedicationByBorrower(borrowerID int) ([]model.Medications, error) {
	var Medication []model.Medications

	// select * from Medication join borrows on Medication.MedicationID = borrows.MedicationID where borrows.UserID = borrowerID
	err := r.db.Joins("join borrows on Medication.MedicationID = borrows.MedicationID").
		Where("borrows.UserID = ?", borrowerID).
		Find(&Medication).Error
	if err != nil {
		log.Printf("GetMedicationByBorrower: Failed to get Medication: %v\n", err)
		return nil, fmt.Errorf("cannot find Medication by borrowerID with error: %v", err)
	}

	return Medication, nil
}

func (r *Repository) GetMedicationByBorrow(borrowID int) (*model.Medications, error) {
	var medication model.Medications

	// Выполнить запрос, выбирая только поля из таблицы books
	err := r.db.Select("Medication.*").Joins("join borrows on Medication.MedicationID = borrows.MedicationID").
		Where("borrows.borrow_id = ?", borrowID).
		Find(&medication).Error
	if err != nil {
		log.Printf("GetMedicationByBorrow: Failed to get Medication: %v\n", err)
		return nil, fmt.Errorf("cannot find Medication by borrowID with error: %v", err)
	}

	return &medication, nil
}

func (r *Repository) UpdateMedication(m *model.Medications) (*model.Medications, error) {
	// update medication set Leaflet = '1 руз 1 шаб', ManufacturerID = 1 where MedicationID = 1
	result := r.db.Model(&m).Updates(&m)
	if result.Error != nil {
		log.Printf("UpdateMedication: Failed to update Medication: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to update Medication: %v\n", result.Error)
	}

	return m, nil
}

func (r *Repository) DeleteMedication(MedicationID int) (int, error) {
	// delete from Medication where MedicationID = Medication_ID returning MedicationID
	result := r.db.Delete(&model.Medications{}, MedicationID)
	if result.Error != nil {
		log.Printf("DeleteMedication: Failed to delete medication: %v\n", result.Error)
		return 0, fmt.Errorf("Failed to delete medication: %v\n", result.Error)
	}

	return MedicationID, nil
}

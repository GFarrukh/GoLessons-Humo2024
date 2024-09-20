package service

import (
	"errors"
	"fmt"
	"pharmacy-system/internal/model"
)

var (
	ErrRecordNotFound     = errors.New("record not found")
	ErrNoMedicationsFound = errors.New("no medications found")
)

func (s *Service) CreateMedication(m *model.Medications) (*model.Medications, error) {
	_, err := s.Repository.GetManufacturerByID(m.ManufacturerID)
	if err != nil {
		if errors.Is(err, ErrRecordNotFound) {
			return nil, fmt.Errorf("manufacturer with id %d doesn't exists", m.ManufacturerID)
		}
		return nil, err
	}

	medications, err := s.Repository.GetMedicationByManufacturer(m.ManufacturerID)
	if err != nil {
		return nil, err
	}

	if len(medications) > 0 {
		for _, medication := range medications {
			if medication.NameMedication == m.NameMedication {
				return nil, fmt.Errorf("the medications whith NameMedication %s and ManufacturerID %d is already exists", m.NameMedication, m.ManufacturerID)
			}
		}
	}

	return s.Repository.AddMedication(m)
}

func (s *Service) ListMedications() ([]model.Medications, error) {
	medications, err := s.Repository.GetMedication()
	if err != nil {
		return nil, err
	}
	if len(medications) == 0 {
		return nil, ErrNoMedicationsFound
	}

	return s.Repository.GetMedication()
}

func (s *Service) FindMedication(id int) (*model.Medications, error) {
	medication, err := s.Repository.GetMedicationByID(id)
	if err != nil {
		return nil, err
	}

	if medication == nil {
		return nil, fmt.Errorf("medications with id %d not found", id)
	}

	return medication, nil
}

func (s *Service) FindMedicationsByManufacturer(ManufacturerID int) ([]model.Medications, error) {
	MedicationByManufacturer, err := s.Repository.GetMedicationByManufacturer(ManufacturerID)
	if err != nil {
		return nil, err
	}
	if len(MedicationByManufacturer) == 0 {
		return nil, fmt.Errorf("no medication with manufacturerID %d", ManufacturerID)
	}

	return s.Repository.GetMedicationByManufacturer(ManufacturerID)
}

func (s *Service) EditMedication(m *model.Medications) (*model.Medications, error) {
	_, err := s.Repository.GetMedicationByID(m.MedicationID)
	if err != nil {
		if errors.Is(err, ErrRecordNotFound) {
			return nil, fmt.Errorf("Medication with id %d not found", m.MedicationID)
		}
		return nil, err
	}

	return s.Repository.UpdateMedication(m)
}

func (s *Service) RemoveMedication(id int) (int, error) {
	_, err := s.Repository.GetMedicationByID(id)
	if err != nil {
		if errors.Is(err, ErrRecordNotFound) {
			return 0, fmt.Errorf("Medication with id %d not found", id)
		}
		return 0, err
	}

	return s.Repository.DeleteMedication(id)
}

func (s *Service) GetMedicationsByManufacturer(ManufacturerID int) ([]model.Medications, error) {
	MedicationsByManufacturer, err := s.Repository.GetMedicationByManufacturer(ManufacturerID)
	if err != nil {
		return nil, err
	}
	if len(MedicationsByManufacturer) == 0 {
		return nil, fmt.Errorf("no medications with ManufacturerID %d", ManufacturerID)
	}

	return s.Repository.GetMedicationByManufacturer(ManufacturerID)
}

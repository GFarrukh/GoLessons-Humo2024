package service

import (
	"errors"
	"fmt"
	"pharmacy-system/internal/model"
)

func (s *Service) CreateManufacturer(m *model.Manufacturer) (*model.Manufacturer, error) {
	_, err := s.Repository.GetManufacturerByNameСompany(m.NameСompany)
	if err != nil && errors.Is(err, ErrNoBooksFound) {
		return nil, err
	} else if !errors.Is(err, ErrNoBooksFound) {
		return nil, fmt.Errorf("author with name %s already exists", m.NameСompany)
	}

	manufacturer, err := s.Repository.AddManufacturer(m)
	if err != nil {
		return nil, err
	}

	return manufacturer, nil
}

func (s *Service) GetManufacturer() ([]model.Manufacturer, error) {
	manufacturer, err := s.Repository.GetManufacturer()
	if err != nil {
		return nil, err
	}

	if len(manufacturer) == 0 {
		return nil, fmt.Errorf("no authors found")
	}

	return manufacturer, nil
}

func (s *Service) GetManufacturerByID(ManufacturerID int) (*model.Manufacturer, error) {
	ManufacturerByID, err := s.Repository.GetManufacturerByID(ManufacturerID)
	if err != nil {
		return nil, err
	}
	if ManufacturerByID == nil {
		return nil, fmt.Errorf("author with id %d not found", ManufacturerID)
	}

	return ManufacturerByID, nil
}

func (s *Service) EditManufacurer(m *model.Manufacturer) (*model.Manufacturer, error) {
	ManufacturerByID, err := s.Repository.GetManufacturerByID(m.ManufacturerID)
	if err != nil {
		return nil, err
	}
	if ManufacturerByID == nil {
		return nil, fmt.Errorf("author with id %d not found", m.ManufacturerID)
	}

	return s.Repository.UpdateManufacturer(m)
}

func (s *Service) DeleteManufacturer(ManufacturerID int) (int, error) {
	ManufacturerByID, err := s.Repository.GetManufacturerByID(ManufacturerID)
	if err != nil {
		return 0, err
	}
	if ManufacturerByID == nil {
		return 0, fmt.Errorf("author with id %d not found", ManufacturerID)
	}

	return s.Repository.DeleteManufacturer(ManufacturerID)
}

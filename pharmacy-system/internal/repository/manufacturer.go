package repository

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"pharmacy-system/internal/model"
)

func (r *Repository) AddManufacturer(m *model.Manufacturer) (*model.Manufacturer, error) {
	// insert into Manufacturer (NameСompany, Leaflet, AddressCompany) values ('Diamed', 'd fjk vdfjk vjkdf v', 'Tajikistan,Dushanbe')
	result := r.db.Create(&m)
	if result.Error != nil {
		log.Printf("CreateManufacturer: Failed to add Manufacturer: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to add Manufacturer: %v\n", result.Error)
	}

	return m, nil
}

func (r *Repository) GetManufacturer() ([]model.Manufacturer, error) {
	var Manufacturer []model.Manufacturer

	// select * from Manufacturer
	result := r.db.Find(&Manufacturer)
	if result.Error != nil {
		log.Printf("GetManufacturer: Failed to get Manufacturer: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to get Manufacturer: %v\n", result.Error)
	}

	return Manufacturer, nil
}

func (r *Repository) GetManufacturerByID(ManufacturerID int) (*model.Manufacturer, error) {
	var Manufacturer model.Manufacturer

	// select * from Manufacturer where ManufacturerID = ManufacturerID
	result := r.db.First(&Manufacturer, ManufacturerID)
	if result.Error == gorm.ErrRecordNotFound {
		log.Printf("GetManufacturerByID: error: %v", result.Error)
		return nil, result.Error
	}
	if result.Error != nil {
		log.Printf("GetManufacturerByID: Failed to get Manufacturer: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to get Manufacturer: %v\n", result.Error)
	}

	return &Manufacturer, nil
}

func (r *Repository) GetManufacturerByNameСompany(NameСompany string) (*model.Manufacturer, error) {
	var Manufacturer model.Manufacturer

	// select * from Manufacturer where NameСompany = NameСompany
	result := r.db.Where("NameСompany = ?", NameСompany).First(&Manufacturer)
	if result.Error != nil {
		log.Printf("GetAuthorByNameСompany: Failed to get Manufacturer: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to get Manufacturer: %v\n", result.Error)
	}

	return &Manufacturer, nil
}

func (r *Repository) UpdateManufacturer(m *model.Manufacturer) (*model.Manufacturer, error) {
	// update Manufacturer set NameCompany = 'Diamed', Leaflet = 'sdckmskldc', AddressCompany = 'Tajikistan' where ManufacturerID = 1
	result := r.db.Model(&m).Updates(&m)
	if result.Error != nil {
		log.Printf("EditManufacturer: Failed to update Manufacturer: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to update Manufacturer: %v\n", result.Error)
	}

	return m, nil
}

func (r *Repository) DeleteManufacturer(ManufacturerID int) (int, error) {
	// delete from Manufacturer where ManufacturerID = ManufacturerID
	result := r.db.Delete(&model.Manufacturer{}, ManufacturerID)
	if result.Error != nil {
		log.Printf("DeleteManufacturer: Failed to delete Manufacturer: %v\n", result.Error)
		return 0, fmt.Errorf("Failed to delete Manufacturer: %v\n", result.Error)
	}

	return ManufacturerID, nil
}

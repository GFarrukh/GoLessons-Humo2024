package service

import "pharmacy-system/internal/model"

type ServiceInterface interface {
	CreateAuthor(a *model.Manufacturer) (*model.Manufacturer, error)
	GetAuthors() ([]model.Manufacturer, error)
	GetAuthorByID(ManufacturerID int) (*model.Manufacturer, error)
	EditAuthor(a *model.Manufacturer) (*model.Manufacturer, error)
	DeleteAuthor(ManufacturerID int) (int, error)
}

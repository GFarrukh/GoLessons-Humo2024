package model

type Manufacturer struct {
	ManufacturerID int    `json:"manufacturer_id,omitempty" gorm:"primaryKey"` //ID Инструкции
	Leaflet        string `json:"leaflet,omitempty"`                           //Инструкция
	NameСompany    string `json:"name_company,omitempty"`                      //Название компании
	AddressCompany string `json:"address_company,omitempty"`                   //Аддрес компании
}

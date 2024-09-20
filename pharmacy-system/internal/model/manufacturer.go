package model

type Manufacturer struct {
	ManufacturerID int    `json:"ManufacturerID,omitempty" gorm:"primaryKey"` //ID Инструкции
	Leaflet        string `json:"biography,omitempty"`                        //Инструкция
	NameСompany    string `json:"NameCompany,omitempty"`                      //Название компании
	AddressCompany string `json:"AddressCompanys,omitempty"`                  //Аддрес компании
}

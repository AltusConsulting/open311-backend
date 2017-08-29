package models

type Address struct {
	AddressString string   `json:"address_string"`
	Location      Location `json:"location"`
	ZipCode       string   `json:"zipcode"`
}

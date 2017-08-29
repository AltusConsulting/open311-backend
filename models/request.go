package models

type Request struct {
	ID                string       `json:"id,omitempty"`
	JurisdictionID    string       `json:"jurisdiction_id"`
	ServiceCode       string       `json:"service_code"`
	Location          Location     `json:"location,omitempty"`
	Attributes        []Attributes `json:"attributes,omitempty"`
	AddressString     string       `json:"address_string,omitempty"`
	AddressID         string       `json:"address_id,omitempty"`
	Email             string       `json:"email"`
	DeviceID          string       `json:"device_id"`
	AccountID         string       `json:"account_id"`
	FirstName         string       `json:"first_name"`
	LastName          string       `json:"last_name"`
	Phone             string       `json:"phone,omitempty"`
	Description       string       `json:"description"`
	MediaURL          string       `json:"media_url"`
	Status            string       `json:"status,omitempty"`
	StatusNotes       string       `json:"status_notes,omitempty"`
	ServiceName       string       `json:"service_name,omitempty"`
	AgencyResponsible string       `json:"agency_responsible,omitempty"`
	ServiceNotice     string       `json:"service_notice,omitempty"`
	RequestedDateTime string       `json:"requested_datetime,omitempty"`
	UpdatedDateTime   string       `json:"updated_datetime,omitempty"`
	ExpectedDateTime  string       `json:"expected_datetime,omitempty"`
	ZipCode           string       `json:"zipcode,omitempty"`
}

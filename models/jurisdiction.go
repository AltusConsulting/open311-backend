package models

type Jurisdiction struct {
	JurisdictionID string `json:"jurisdiction_id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	Country        string `json:"country"`
	Area           string `json:"area"`
}

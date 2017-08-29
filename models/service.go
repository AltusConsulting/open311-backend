package models

type Service struct {
	ID             string `json:"id, omitempty"`
	JurisdictionID string `json:"jurisdiction_id"`
	ServiceName    string `json:"service_name"`
	Description    string `json:"description"`
	Metadata       bool   `json:"metadata"`
	Type           string `json:"type"`
	Keywords       string `json:"keywords"`
	Group          string `json:"group"`
}

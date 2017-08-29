package models

type ServiceDefinition struct {
	JurisdictionID       string            `json:"jurisdiction_id"`
	ServiceCode          string            `json:"service_code"`
	Variable             bool              `json:"variable"`
	DataType             string            `json:"datatype"`
	Required             bool              `json:"required"`
	DataTypeDescription  string            `json:"datatype_description"`
	Order                int               `json:"order"`
	AttributeDescription string            `json:"attribute_description"`
	Value                map[string]string `json:"value"`
}

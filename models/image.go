package models

type Image struct {
	BucketName string `json:"bucket_name"`
	ObjectName string `json:"object_name"`
	Image      string `json:"image"`
}

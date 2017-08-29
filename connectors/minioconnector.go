package connectors

import (
	minio "github.com/minio/minio-go"
	"github.com/spf13/viper"
)

func CreateMinioClient() (*minio.Client, error) {
	endpoint := viper.GetString("minio.host")
	accesKeyID := viper.GetString("minio.accessKeyID")
	secretAccessKey := viper.GetString("minio.secretAccessKey")
	useSSL := false

	minioClient, err := minio.New(endpoint, accesKeyID, secretAccessKey, useSSL)

	return minioClient, err
}

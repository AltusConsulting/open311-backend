package views

import (
	"encoding/base64"
	"errors"
	"net/http"
	"os"
	"time"

	"../models"

	"fmt"

	"net/url"

	"../connectors"

	log "github.com/Sirupsen/logrus"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

//Get image
func PutRequestImage(c *gin.Context) {

	var image models.Image
	err := c.BindJSON(&image)
	if err != nil {
		fmt.Println("Bad json")
		c.AbortWithError(http.StatusBadRequest, errors.New("Bad Json"))
		return
	}

	//log.Info(fmt.Sprintf("%v", image))

	//initialize s3 client
	s3client, err := connectors.CreateS3Client(viper.GetString("s3.region"))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	decoded, err := base64.StdEncoding.DecodeString(image.Image)
	if err != nil {
		fmt.Println("decode error:", err)
		return
	}
	//fmt.Println("decoded" + string(decoded))

	file, err := os.Create(image.ObjectName)
	if err != nil {
		fmt.Println("error while create")
		fmt.Println(err)
	}
	defer file.Close()

	n2, err := file.Write(decoded)
	if err != nil {
		fmt.Println("error while write")
		fmt.Println(err)
	}
	fmt.Printf("wrote %d bytes\n", n2)

	file2, err := os.Open(image.ObjectName)
	if err != nil {
		fmt.Println("error while open")
		fmt.Println(err)
		return
	}
	defer file.Close()

	_, err = s3client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(image.BucketName),
		Key:    aws.String(image.ObjectName),
		Body:   file2,
	})
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	log.Printf("Successfully uploaded %s\n", image.ObjectName)

	c.JSON(http.StatusOK, gin.H{"code": 200, "description": "Image saved"})

}

// Get image URL
func GetRequestImage(c *gin.Context) {
	bucketName := c.Query("bucketName")
	objectName := c.Query("objectName")

	if bucketName == "" || objectName == "" {
		c.AbortWithError(http.StatusBadRequest, errors.New("Missing query"))
		return
	}

	// Initialize s3 client object
	s3Client, err := connectors.CreateS3Client(viper.GetString("s3.region"))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	reqParams := make(url.Values)
	reqParams.Set("response-content-disposition", "attachment; filename=\""+objectName+"\"")

	result, err := s3Client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectName),
	})

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchKey:
				c.JSON(http.StatusBadRequest, gin.H{"code": 404, "description": "The specified key does not exist"})
			default:
				c.JSON(http.StatusBadRequest, gin.H{"code": 400, "description": aerr.Error()})
			}
		}
		fmt.Println(result)
		return
	}

	req, _ := s3Client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectName),
	})

	// Generates a presigned url which expires in a day.
	presignedURL, err := req.Presign(time.Second * 168 * 60 * 60) // 50 years
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "description": err})
		return
	}
	fmt.Println(presignedURL)

	c.JSON(http.StatusOK, gin.H{"code": 200, "url": presignedURL})
}

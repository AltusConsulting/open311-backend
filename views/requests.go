package views

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"../connectors"
	"../models"

	elastic "gopkg.in/olivere/elastic.v5"

	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

func ServiceRequest(c *gin.Context) {
	var request models.Request
	err := c.BindJSON(&request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("Bad Json"))
		return
	}
	log.Info(fmt.Sprintf("%v", request))

	ctx := context.Background()
	client, err := connectors.Create311Client()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, errors.New("failed to create client"))
		return
	}

	response, err := client.Index().
		Index("request").
		Type("request").
		BodyJson(request).
		Refresh("true").
		Do(ctx)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, errors.New("failed to send json"))
		return
	}
	c.Header("Location", fmt.Sprintf("%s/%s", c.Request.URL, response.Id))
	c.JSON(http.StatusOK, gin.H{"message": response.Created})
}

//Get requests
func GetServiceRequests(c *gin.Context) {
	ctx := context.Background()

	client, err := connectors.Create311Client()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	result, err := client.Search().
		Index("request").
		Type("request").
		From(0).Size(1000).
		Do(ctx)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	if result.Hits.TotalHits > 0 {
		var response []interface{}

		for _, source := range result.Hits.Hits {
			var item models.Request
			err := json.Unmarshal(*source.Source, &item)
			if err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
			}
			//log.Info(item)
			item.ID = source.Id
			response = append(response, item)

		}
		c.JSON(http.StatusOK, response)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "No items"})
	}
}

func GetServiceRequestByID(c *gin.Context) {
	serviceRequestID := c.Param("service_request_id")

	ctx := context.Background()

	client, err := connectors.Create311Client()
	if err != nil {
		panic(err)
	}

	termQuery := elastic.NewTermQuery("_id", serviceRequestID)
	result, err := client.Search().
		Index("request").
		Type("request").
		Query(termQuery).
		Do(ctx)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	fmt.Println(result.Hits.TotalHits)
	if result.Hits.TotalHits > 0 {
		var item models.Request
		for _, source := range result.Hits.Hits {
			err := json.Unmarshal(*source.Source, &item)
			if err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
			}
		}
		c.JSON(http.StatusOK, item)
	} else {
		c.AbortWithError(http.StatusNotFound, errors.New("Item not found"))
	}

}

package views

import (
	"context"
	"net/http"

	"../models"

	"encoding/json"

	"fmt"

	"errors"

	"../connectors"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	elastic "gopkg.in/olivere/elastic.v5"
)

func GetServiceList(c *gin.Context) {
	ctx := context.Background()

	client, err := connectors.Create311Client()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	searchResult, err := client.Search().
		Index("service").
		Type("service").
		Do(ctx)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	if searchResult.Hits.TotalHits > 0 {

		var response []interface{}

		for _, source := range searchResult.Hits.Hits {
			var item models.Service
			err := json.Unmarshal(*source.Source, &item)
			if err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
			}
			item.ID = source.Id
			response = append(response, item)
		}
		c.JSON(http.StatusOK, response)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "No items"})
	}
}

func GetServiceDefinition(c *gin.Context) {
	serviceCode := c.Param("service_code")
	log.Debug(fmt.Sprintf("service code: %v", serviceCode))
	ctx := context.Background()
	client, err := connectors.Create311Client()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	termQuery := elastic.NewTermQuery("service_code", serviceCode)
	searchResult, err := client.Search().
		Index("service").
		Type("service_definition").
		Query(termQuery).
		Do(ctx)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	if searchResult.Hits.TotalHits > 0 {
		var item models.ServiceDefinition
		for _, source := range searchResult.Hits.Hits {
			err = json.Unmarshal(*source.Source, &item)
			if err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
			}
		}
		c.JSON(http.StatusOK, item)

	} else {
		c.AbortWithError(http.StatusNotFound, errors.New("item not found"))
	}
}

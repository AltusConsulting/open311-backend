package views

import (
	"testing"

	"../models"

	"net/http"

	"encoding/json"
	"log"

	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/h2non/gock.v1"
)

func TestGetServiceList(t *testing.T) {
	defer gock.Off()

	Convey("Make a request to get a list of services type", t, func() {
		var service models.Service
		service.Description = "description"
		service.Group = "group"
		service.ID = "ID"
		service.JurisdictionID = "Jurisdiction ID"
		service.Keywords = "keywords"
		service.Metadata = false
		service.ServiceName = "service name"
		service.Type = "type"
		serviceArray := []models.Service{service, service}
		Convey("make a GET to a mock database", func() {
			gock.New("http://localhost:8080/api/311/v1").
				Get("/services").
				Reply(200).
				JSON(serviceArray)

			res, _ := http.Get("http://localhost:8080/api/311/v1/services")
			So(res.StatusCode, ShouldEqual, http.StatusOK)
			var result []models.Service
			if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
				log.Fatalln(err)
			}
			So(result, ShouldResemble, serviceArray)
			So(result, ShouldHaveSameTypeAs, serviceArray)
		})

	})
}

func TestGetServiceCode(t *testing.T) {
	defer gock.Off()

	Convey("make a request to get a service definition info", t, func() {
		var serviceDefinition models.ServiceDefinition
		serviceDefinition.AttributeDescription = "attribute description"
		serviceDefinition.DataTypeDescription = "data type description"
		serviceDefinition.DataType = "data type"
		serviceDefinition.JurisdictionID = "jurisdiction ID"
		serviceDefinition.Order = 1
		serviceDefinition.Required = false
		serviceDefinition.ServiceCode = "service code"
		m := make(map[string]string)
		m["code"] = "value"
		serviceDefinition.Value = m
		serviceDefinition.Variable = false
		Convey("make a GET to a mock database", func() {
			gock.New("http://localhost:8080/api/311/v1").
				Get("/services/1").
				Reply(200).
				JSON(serviceDefinition)

			res, _ := http.Get("http://localhost:8080/api/311/v1/services/1")
			So(res.StatusCode, ShouldEqual, http.StatusOK)
			var result models.ServiceDefinition
			if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
				log.Fatalln(err)
			}

			So(result, ShouldHaveSameTypeAs, serviceDefinition)
			So(result, ShouldResemble, serviceDefinition)
		})
	})
}

package views

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"../models"

	"bytes"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/h2non/gock.v1"
)

func TestRequest(t *testing.T) {
	defer gock.Off()

	Convey("Make a http request to insert a new report request", t, func() {
		var request = getMockRequest()

		Convey("make a POST to a mock database", func() {
			gock.New("http://localhost:8080/api/311/v1").
				Post("/requests").
				Reply(200).
				JSON(gin.H{"created": true}).
				Header.Set("Location", "http://localhost:8080/api/311/v1/requests/1")

			body := bytes.NewBuffer([]byte(fmt.Sprintf("%v", request)))
			res, _ := http.Post("http://localhost:8080/api/311/v1/requests", "application/json", body)
			So(res.StatusCode, ShouldEqual, http.StatusOK)
			So(res.Header.Get("Location"), ShouldEqual, "http://localhost:8080/api/311/v1/requests/1")
			var result map[string]interface{}
			if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
				log.Fatalln(err)
			}

			So(result["created"], ShouldEqual, true)
		})

		Convey("make a POST to a mock database with a bad json", func() {
			badJson := getBadJsonRequest()
			gock.New("http://localhost:8080/api/311/v1").
				Post("/requests").
				Reply(http.StatusBadRequest).
				JSON(gin.H{"message": "Bad json"})

			body := bytes.NewBuffer([]byte(fmt.Sprintf("%v", badJson)))
			res, _ := http.Post("http://localhost:8080/api/311/v1/requests", "application/json", body)
			So(res.StatusCode, ShouldEqual, http.StatusBadRequest)
			So(res.Header.Get("Location"), ShouldBeEmpty)

			var result map[string]interface{}
			if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
				log.Fatalln(err)
			}

			So(result["message"], ShouldEqual, "Bad json")
		})

	})
}

func TestGetRequestsList(t *testing.T) {
	defer gock.Off()

	Convey("make a http request to get a list of report requests", t, func() {
		var request = getMockResponse()

		requestArray := []models.Request{request, request}

		Convey("make a GET to a mock database", func() {
			gock.New("http://localhost:8080/api/311/v1").
				Get("/requests").
				Reply(200).
				JSON(requestArray)

			res, _ := http.Get("http://localhost:8080/api/311/v1/requests")
			So(res.StatusCode, ShouldEqual, http.StatusOK)

			var result []models.Request
			if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
				log.Fatalln(err)
			}

			So(result, ShouldResemble, requestArray)
			So(result, ShouldHaveSameTypeAs, requestArray)
		})
	})
}

func TestGetRequestByID(t *testing.T) {
	defer gock.Off()

	Convey("make a request to get a report request info by its ID", t, func() {
		request := getMockResponse()

		Convey("make a GET to a mock database", func() {
			gock.New("http://localhost:8080/api/311/v1").
				Get("/requests/1").
				Reply(200).
				JSON(request)

			res, _ := http.Get("http://localhost:8080/api/311/v1/requests/1")
			So(res.StatusCode, ShouldEqual, http.StatusOK)

			var result models.Request
			if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
				log.Fatalln(err)
			}

			So(result, ShouldHaveSameTypeAs, request)
			So(result, ShouldResemble, request)
		})
	})
}

func getMockRequest() models.Request {
	var request models.Request
	request.JurisdictionID = "jurisdiction ID"
	request.ServiceCode = "service code"
	var lat models.Location
	lat.Lat = 10.00
	lat.Lon = 10.00
	request.Location = lat
	request.AddressString = "address string"
	request.AddressID = "address ID"
	request.Email = "email@mail.com"
	request.DeviceID = "deviceID"
	request.AccountID = "accountID"
	request.FirstName = "first name"
	request.LastName = "last name"
	request.Phone = "88888888"
	request.Description = "description"
	request.MediaURL = "media url"

	return request
}

func getBadJsonRequest() models.Request {
	var request models.Request
	request.AddressString = "address string"
	request.AddressID = "address ID"
	request.Email = "email@mail.com"
	request.DeviceID = "deviceID"
	request.AccountID = "accountID"
	request.FirstName = "first name"
	request.LastName = "last name"
	request.Phone = "88888888"
	request.Description = "description"

	return request
}

func getMockResponse() models.Request {
	var request = getMockRequest()
	request.Status = "status"
	request.StatusNotes = "status notes"
	request.ServiceName = "service name"
	request.AgencyResponsible = "agency responsible"
	request.ServiceNotice = "service notice"
	request.RequestedDateTime = "2010-10-10"
	request.UpdatedDateTime = "2010-10-10"
	request.ExpectedDateTime = "2010-10-10"
	request.ZipCode = "10312"

	return request
}

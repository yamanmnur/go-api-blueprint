package controllers

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"

	base_responses "go-crud-api/pkg/common/responses"
	data "go-crud-api/pkg/data/cats/data"
	mocks "go-crud-api/pkg/mocks/cats"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
  Actual test functions
*/

// TestSomething is an example of how to use our test object to
// make assertions about some target code we are testing.
func TestCatsData(t *testing.T) {

	// create an instance of our test object
	catService := new(mocks.ICatService)

	// setup expectations
	mockData := data.CatsData{
		Id:     1,
		Name:   "Agni",
		Race:   "Persia",
		Age:    20,
		Gender: "M",
	}
	catService.On("GetCatsById", 1).Return(mockData, nil)

	catController := CatController{catService}

	// call the code we are testing
	router := gin.Default()
	router.GET("/cat/:id", func(c *gin.Context) {
		catController.GetById(c)
	})

	req := httptest.NewRequest("GET", "/cat/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	fmt.Print("BODY RESPONSE")
	fmt.Print(w.Body)

	responseBody := base_responses.GenericResponse{
		Data: data.CatsData{},
	}
	// Decode the entire JSON response
	json.NewDecoder(w.Body).Decode(&responseBody)
	actualResponseData := data.CatsData{}
	err := mapstructure.Decode(responseBody.Data, &actualResponseData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Extract the "data" field from the response
	expectedResult := data.CatsData{
		Id:     1,
		Name:   "Agni",
		Race:   "Persia",
		Age:    20,
		Gender: "M",
	}
	// assert that the expectations were met
	assert.Equal(t, expectedResult, actualResponseData)
}

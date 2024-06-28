package tests

import (
	"WeekendPOS/app/entity"
	"WeekendPOS/app/model"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateEmployeeSuccess(t *testing.T) {
	TestLoginSuccess(t)

	user := new(entity.User)
	err := db.Where("username = ?", "johndoe").First(user).Error
	assert.Nil(t, err)

	requestBody := model.CreateEmployeeRequest{
		Name:     "Thomas",
		Email:    "thomas@gmail.com",
		Username: "thomas",
		Password: "rahasia",
		Phone:    "081234567890",
		Address:  "Jl. Jendral Sudirman No. 1",
	}

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/api/employee", strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.EmployeeResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, requestBody.Name, responseBody.Data.Name)
	assert.Equal(t, requestBody.Email, responseBody.Data.Email)
	assert.Equal(t, requestBody.Username, responseBody.Data.Username)
	assert.Equal(t, requestBody.Phone, responseBody.Data.Phone)
	assert.Equal(t, requestBody.Address, responseBody.Data.Address)
}

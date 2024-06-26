package tests

import (
	"WeekendPOS/app/model"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateProductSuccess(t *testing.T) {
	TestCreateCategorySuccess(t)
	user := GetFirstUser(t)
	category := GetFirstCategory(t, user)

	requestBody := model.CreateProductRequest{
		SKU:        "SKU-001",
		Name:       "Product 1",
		CategoryID: category.ID,
		BuyPrice:   10000,
		SellPrice:  15000,
		Stock:      10,
	}

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/api/product/", strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.ProductResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, requestBody.SKU, responseBody.Data.SKU)
	assert.Equal(t, requestBody.Name, responseBody.Data.Name)
	assert.Equal(t, category.Name, responseBody.Data.Category)
}

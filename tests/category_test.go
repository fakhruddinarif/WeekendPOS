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

func TestCreateCategorySuccess(t *testing.T) {
	ClearAll()
	TestLoginSuccess(t)

	user := new(entity.User)
	err := db.Where("username = ?", "johndoe").First(user).Error
	assert.Nil(t, err)

	requestBody := model.CreateCategoryRequest{
		Name: "Food",
	}

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/api/category", strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.CategoryResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, requestBody.Name, responseBody.Data.Name)
}

func TestCreateCategoryFailed(t *testing.T) {
	ClearAll()
	TestLoginSuccess(t)

	user := new(entity.User)
	err := db.Where("username = ?", "johndoe").First(user).Error
	assert.Nil(t, err)

	requestBody := model.CreateCategoryRequest{
		Name: "",
	}

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/api/category", strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.CategoryResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	assert.NotNil(t, responseBody.Errors)
}

func TestGetCategorySuccess(t *testing.T) {
	TestCreateCategorySuccess(t)

	user := new(entity.User)
	err := db.Where("username = ?", "johndoe").First(user).Error
	assert.Nil(t, err)

	category := new(entity.Category)
	err = db.Where("user_id = ?", user.ID).First(category).Error
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodGet, "/api/category/"+category.ID, nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.CategoryResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, category.ID, responseBody.Data.ID)
	assert.Equal(t, category.Name, responseBody.Data.Name)
}

func TestGetCategoryFailed(t *testing.T) {
	TestCreateCategorySuccess(t)

	user := new(entity.User)
	err := db.Where("username = ?", "johndoe").First(user).Error
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodGet, "/api/category/0", nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.CategoryResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusNotFound, response.StatusCode)
	assert.NotNil(t, responseBody.Errors)
}

func TestUpdateCategorySuccess(t *testing.T) {
	TestCreateCategorySuccess(t)

	user := new(entity.User)
	err := db.Where("username = ?", "johndoe").First(user).Error
	assert.Nil(t, err)

	category := new(entity.Category)
	err = db.Where("user_id = ?", user.ID).First(category).Error
	assert.Nil(t, err)

	requestBody := model.UpdateCategoryRequest{
		Name: "Drink",
	}

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPut, "/api/category/"+category.ID, strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.CategoryResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, requestBody.Name, responseBody.Data.Name)
	assert.NotNil(t, responseBody.Data.ID)
	assert.NotNil(t, responseBody.Data.CreatedAt)
	assert.NotNil(t, responseBody.Data.UpdatedAt)
}

func TestUpdateCategoryFailed(t *testing.T) {
	TestCreateCategorySuccess(t)

	user := new(entity.User)
	err := db.Where("username = ?", "johndoe").First(user).Error
	assert.Nil(t, err)

	category := new(entity.Category)
	err = db.Where("user_id = ?", user.ID).First(category).Error
	assert.Nil(t, err)

	requestBody := model.UpdateCategoryRequest{
		Name: "",
	}

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPut, "/api/category/"+category.ID, strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.CategoryResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	assert.NotNil(t, responseBody.Errors)
}

func TestUpdateCategoryNotFound(t *testing.T) {
	TestCreateCategorySuccess(t)

	user := new(entity.User)
	err := db.Where("username = ?", "johndoe").First(user).Error
	assert.Nil(t, err)

	requestBody := model.UpdateCategoryRequest{
		Name: "Drink",
	}

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPut, "/api/category/0", strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.CategoryResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusNotFound, response.StatusCode)
	assert.NotNil(t, responseBody.Errors)
}

func TestDeleteCategorySuccess(t *testing.T) {
	TestCreateCategorySuccess(t)

	user := new(entity.User)
	err := db.Where("username = ?", "johndoe").First(user).Error
	assert.Nil(t, err)

	category := new(entity.Category)
	err = db.Where("user_id = ?", user.ID).First(category).Error
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodDelete, "/api/category/"+category.ID, nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[bool])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, true, responseBody.Data)
}

func TestDeleteCategoryFailed(t *testing.T) {
	TestCreateCategorySuccess(t)

	user := new(entity.User)
	err := db.Where("username = ?", "johndoe").First(user).Error
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodDelete, "/api/category/0", nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[bool])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusNotFound, response.StatusCode)
	assert.NotNil(t, responseBody.Errors)
}

func TestSearchCategorySuccess(t *testing.T) {
	TestLoginSuccess(t)

	user := new(entity.User)
	err := db.Where("username = ?", "johndoe").First(user).Error
	assert.Nil(t, err)

	CreateCategories(user, 20)

	request := httptest.NewRequest(http.MethodGet, "/api/category/", nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[[]model.CategoryResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, 10, len(responseBody.Data))
	assert.Equal(t, int64(20), responseBody.Paging.TotalItem)
	assert.Equal(t, int64(2), responseBody.Paging.TotalPage)
	assert.Equal(t, 1, responseBody.Paging.Page)
	assert.Equal(t, 10, responseBody.Paging.Size)
}

func TestSearchCategoryWithPagination(t *testing.T) {
	TestLoginSuccess(t)

	user := new(entity.User)
	err := db.Where("username = ?", "johndoe").First(user).Error
	assert.Nil(t, err)

	CreateCategories(user, 20)

	request := httptest.NewRequest(http.MethodGet, "/api/category/?page=2&size=5", nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[[]model.CategoryResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, 5, len(responseBody.Data))
	assert.Equal(t, int64(20), responseBody.Paging.TotalItem)
	assert.Equal(t, int64(4), responseBody.Paging.TotalPage)
	assert.Equal(t, 2, responseBody.Paging.Page)
	assert.Equal(t, 5, responseBody.Paging.Size)
}

func TestSearchCategoryWithFilter(t *testing.T) {
	TestLoginSuccess(t)

	user := new(entity.User)
	err := db.Where("username = ?", "johndoe").First(user).Error
	assert.Nil(t, err)

	CreateCategories(user, 20)

	request := httptest.NewRequest(http.MethodGet, "/api/category/?name=Category 1", nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[[]model.CategoryResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, 1, len(responseBody.Data))
	assert.Equal(t, int64(1), responseBody.Paging.TotalItem)
	assert.Equal(t, int64(1), responseBody.Paging.TotalPage)
	assert.Equal(t, 1, responseBody.Paging.Page)
	assert.Equal(t, 10, responseBody.Paging.Size)
}

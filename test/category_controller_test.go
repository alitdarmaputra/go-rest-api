package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/alitdarmaputra/belajar-golang-rest-api/app"
	"github.com/alitdarmaputra/belajar-golang-rest-api/controller"
	"github.com/alitdarmaputra/belajar-golang-rest-api/exception"
	"github.com/alitdarmaputra/belajar-golang-rest-api/helper"
	"github.com/alitdarmaputra/belajar-golang-rest-api/middleware"
	"github.com/alitdarmaputra/belajar-golang-rest-api/model/domain"
	"github.com/alitdarmaputra/belajar-golang-rest-api/repository"
	"github.com/alitdarmaputra/belajar-golang-rest-api/service"
	"github.com/go-playground/assert/v2"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open(
		"mysql",
		"root:05042002@tcp(localhost:3306)/belajar_golang_restful_api_test",
	)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Second)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupRouter(db *sql.DB) http.Handler {
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)
	router.PanicHandler = exception.ErrorHandler

	return middleware.NewAuthMiddleware(router)
}

func truncateCategory(db *sql.DB) {
	db.Exec("TRUNCATE app_category")
}

func TestCreateCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	router := setupRouter(db)

	requestBody := strings.NewReader(`{ "name": "Gadget" }`)
	request := httptest.NewRequest(
		http.MethodPost,
		"http://localhost:3000/api/categories",
		requestBody,
	)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, _ := ioutil.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusCreated, int(responseBody["code"].(float64)))
	assert.Equal(t, "Created", responseBody["status"])
	assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["name"])
}

func TestCreateCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	router := setupRouter(db)

	requestBody := strings.NewReader(`{ "name": "" }`)
	request := httptest.NewRequest(
		http.MethodPost,
		"http://localhost:3000/api/categories",
		requestBody,
	)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	body, _ := ioutil.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusBadRequest, int(responseBody["code"].(float64)))
	assert.Equal(t, "Bad request", responseBody["status"])
}

func TestUpdateCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{ "name": "Gadget 2" }`)
	request := httptest.NewRequest(
		http.MethodPut,
		"http://localhost:3000/api/categories/"+strconv.Itoa(int(category.Id)),
		requestBody,
	)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, _ := ioutil.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "Gadget 2", responseBody["data"].(map[string]interface{})["name"])
}

func TestUpdateCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{ "name": "" }`)
	request := httptest.NewRequest(
		http.MethodPut,
		"http://localhost:3000/api/categories/"+strconv.Itoa(int(category.Id)),
		requestBody,
	)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	body, _ := ioutil.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusBadRequest, int(responseBody["code"].(float64)))
	assert.Equal(t, "Bad request", responseBody["status"])
}

func TestGetCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(
		http.MethodGet,
		"http://localhost:3000/api/categories/"+strconv.Itoa(int(category.Id)),
		nil,
	)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, _ := ioutil.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["name"])
}

func TestGetCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	router := setupRouter(db)

	request := httptest.NewRequest(
		http.MethodGet,
		"http://localhost:3000/api/categories/10",
		nil,
	)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusNotFound, response.StatusCode)

	body, _ := ioutil.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusNotFound, int(responseBody["code"].(float64)))
	assert.Equal(t, "Not found", responseBody["status"])
}

func TestDeleteCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(
		http.MethodDelete,
		"http://localhost:3000/api/categories/"+strconv.Itoa(int(category.Id)),
		nil,
	)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, _ := ioutil.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
}

func TestDeleteCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	router := setupRouter(db)

	request := httptest.NewRequest(
		http.MethodDelete,
		"http://localhost:3000/api/categories/1",
		nil,
	)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusNotFound, response.StatusCode)

	body, _ := ioutil.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusNotFound, int(responseBody["code"].(float64)))
	assert.Equal(t, "Not found", responseBody["status"])
}

func TestListCategoriesSuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "Laptop",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(
		http.MethodGet,
		"http://localhost:3000/api/categories",
		nil,
	)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, _ := ioutil.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])

	var categories []interface{} = responseBody["data"].([]interface{})

	assert.Equal(
		t,
		"Gadget",
		categories[0].(map[string]interface{})["name"],
	)
	assert.Equal(
		t,
		"Laptop",
		categories[1].(map[string]interface{})["name"],
	)
}

func TestUnauthorized(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "Laptop",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(
		http.MethodGet,
		"http://localhost:3000/api/categories",
		nil,
	)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusUnauthorized, response.StatusCode)

	body, _ := ioutil.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusUnauthorized, int(responseBody["code"].(float64)))
	assert.Equal(t, "Unauthorize", responseBody["status"])
}

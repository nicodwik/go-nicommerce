package controllers

import (
	"go-nicommerce/config"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func InitEcho() *echo.Echo {
	config.InitDB()
	e := echo.New()
	return e
}

func TestGetUserController(t *testing.T) {
	var testCases = []struct {
		name                 string
		path                 string
		expectStatus         int
		expectBodyStartsWith string
	}{
		{
			name:                 "success",
			path:                 "/users",
			expectBodyStartsWith: "{\"status\":\"success\",\"users\":[",
			expectStatus:         http.StatusOK,
		},
	}

	e := InitEcho()
	// localhost:8000/api/v1/member/users
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		// Assertions
		if assert.NoError(t, GetUserController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()
			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
		}
	}
}

func TestGetUserByIDController(t *testing.T) {
	var testCases = []struct {
		name                 string
		path                 string
		expectStatus         int
		expectBodyStartsWith string
	}{
		{
			name:                 "success",
			path:                 "/users/1",
			expectBodyStartsWith: "{\"status\":\"success\",\"users\":[",
			expectStatus:         http.StatusOK,
		},
	}

	e := InitEcho()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {

		c.SetPath(testCase.path)

		// Assertions
		if assert.NoError(t, GetUserController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()
			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
		}
	}
}

func TestInsertUserController(t *testing.T) {
	var testCases = []struct {
		name                 string
		path                 string
		expectStatus         int
		expectBodyStartsWith string
	}{
		{
			name:                 "success",
			path:                 "/users",
			expectBodyStartsWith: "{\"data\":{",
			expectStatus:         http.StatusOK,
		},
	}

	e := InitEcho()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {

		c.SetPath(testCase.path)

		// Assertions
		if assert.NoError(t, InsertUserController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()
			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
		}
	}
}

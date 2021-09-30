package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProductsByStoreIDController(t *testing.T) {
	var testCases = []struct {
		name                 string
		path                 string
		expectStatus         int
		expectBodyStartsWith string
	}{
		{
			name:                 "success",
			path:                 "/product/1",
			expectBodyStartsWith: "{\"data\":[",
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
		if assert.NoError(t, GetProductsByStoreIDController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()
			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
		}
	}
}

func TestInsertProductByCategoryIDController(t *testing.T) {
	var testCases = []struct {
		name                 string
		path                 string
		expectStatus         int
		expectBodyStartsWith string
	}{
		{
			name:                 "success",
			path:                 "/product",
			expectBodyStartsWith: "{\"data\":{",
			expectStatus:         http.StatusOK,
		},
	}

	e := InitEcho()

	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetParamNames("category_id")
		c.SetParamValues("1")
		c.SetPath(testCase.path)

		// Assertions
		if assert.NoError(t, InsertProductByCategoryIDController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()
			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
		}
	}
}

package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetShipmentOptionsByStoreIDController(t *testing.T) {
	var testCases = []struct {
		name                 string
		path                 string
		expectStatus         int
		expectBodyStartsWith string
	}{
		{
			name:                 "success",
			path:                 "/shipment-option/1",
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
		if assert.NoError(t, GetShipmentOptionsByStoreIDController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()
			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
		}
	}
}

func TestInsertShipmentOptionByStoreIDController(t *testing.T) {
	var testCases = []struct {
		name                 string
		path                 string
		expectStatus         int
		expectBodyStartsWith string
	}{
		{
			name:                 "success",
			path:                 "/shipment-option",
			expectBodyStartsWith: "{\"data\":{",
			expectStatus:         http.StatusOK,
		},
	}

	e := InitEcho()

	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetParamNames("store_id")
		c.SetParamValues("1")
		c.SetPath(testCase.path)

		// Assertions
		if assert.NoError(t, InsertShipmentOptionController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()
			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
		}
	}
}

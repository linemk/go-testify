package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestMainHandlerWhenMethodIsRight(t *testing.T) {
	method := "GET"
	req, err := http.NewRequest(method, "/cafe?count=4&city=moscow", nil)
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
		return
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(recorder, req)

	assert.Equal(t, recorder.Code, http.StatusOK)
	assert.NotEmpty(t, recorder.Body.String())

}

func TestMainHandlerWhenCityWrong(t *testing.T) {
	city := "Petrograd"
	actualAnswer := "wrong city value"
	req, err := http.NewRequest("GET", "/cafe?count=4&city="+city, nil)
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
		return
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	assert.Equal(t, actualAnswer, recorder.Body.String())
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req, err := http.NewRequest("GET", "/cafe?count="+strconv.Itoa(totalCount)+"&city=moscow", nil)
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
		return
	}
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.NotEmpty(t, responseRecorder.Body.String())
	assert.Len(t, strings.Split(responseRecorder.Body.String(), ","), totalCount)
}

package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomePage(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()
	HomePage(res, req)
	//checking if the status code of response was
	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, res.Code)

	}

	expected := "Hello this is the homepage"
	if res.Body.String() != expected {
		t.Errorf("Expected body %s, got %s", res.Body.String(), expected)
	}
}
func TestAboutPage(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/about", nil)
	res := httptest.NewRecorder()
	AboutPage(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, res.Code)

	}
	expected := "Hello this is the about page"
	if res.Body.String() != expected {
		t.Errorf("Expected body %s, got %s", res.Body.String(), expected)
	}
}

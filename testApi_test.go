package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllThreads(t *testing.T) {
	req, err := http.NewRequest("GET", "/get_all_threads", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetEntries)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	} else {
		fmt.Printf("handler returned wrong status code: got %v want %v and body %v", status, http.StatusOK, rr.Body.String())
	}
}

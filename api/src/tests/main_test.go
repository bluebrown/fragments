package main

import (
	"net/http"
	"testing"
)

func TestRequest(t *testing.T) {
	_, err := http.Get("http://localhost:8080/query")
	if err != nil {
		t.Error(err)
	}
}

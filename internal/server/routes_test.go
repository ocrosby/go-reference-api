package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetHandler(t *testing.T) {
	mux := SetupRoutes()

	req, err := http.NewRequest("GET", "/get", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status %v; got %v", http.StatusOK, rr.Code)
	}
}

func TestPutHandler(t *testing.T) {
	mux := SetupRoutes()

	req, err := http.NewRequest("PUT", "/put", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status %v; got %v", http.StatusOK, rr.Code)
	}
}

func TestPostHandler(t *testing.T) {
	mux := SetupRoutes()

	req, err := http.NewRequest("POST", "/post", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status %v; got %v", http.StatusOK, rr.Code)
	}
}

func TestDeleteHandler(t *testing.T) {
	mux := SetupRoutes()

	req, err := http.NewRequest("DELETE", "/delete", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status %v; got %v", http.StatusOK, rr.Code)
	}
}

func TestCreateUserHandler(t *testing.T) {
	mux := SetupRoutes()

	req, err := http.NewRequest("POST", "/user", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status %v; got %v", http.StatusOK, rr.Code)
	}
}

func TestLivenessProbeHandler(t *testing.T) {
	mux := SetupRoutes()

	req, err := http.NewRequest("GET", "/healthz", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status %v; got %v", http.StatusOK, rr.Code)
	}

	expected := "{\"status\":\"ok\"}\n"
	if rr.Body.String() != expected {
		t.Errorf("expected body %v; got %v", expected, rr.Body.String())
	}
}

func TestReadinessProbeHandler(t *testing.T) {
	mux := SetupRoutes()

	req, err := http.NewRequest("GET", "/readyz", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status %v; got %v", http.StatusOK, rr.Code)
	}

	expected := "{\"status\":\"ok\"}\n"
	if rr.Body.String() != expected {
		t.Errorf("expected body %v; got %v", expected, rr.Body.String())
	}
}

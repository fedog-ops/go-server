package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"acme/db"
	"encoding/json"
	"reflect"
)

func TestRootHandler(t *testing.T) {
    req, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatal(err)
    }

    // Create a new response recorder
    rr := httptest.NewRecorder()

	// Create a handler
	handler := http.HandlerFunc(rootHandler)

	//ACT
	// Serve the request
	handler.ServeHTTP(rr, req)

	    //ASSERT
    // Check the status code
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    // Check the response body
    expected := "Hello, World!"
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
    }
}

func TestRootHandlerWithServer(t *testing.T) {
    //ARRANGE
	// Create a new server with the handler
    server := httptest.NewServer(http.HandlerFunc(rootHandler))
    defer server.Close()

	//ACT
    // Send a GET request to the server
    resp, err := http.Get(server.URL + "/")
    if err != nil {
        t.Fatalf("Failed to send GET request: %v", err)
    }
    defer resp.Body.Close()

	//check status code
	if status := resp.StatusCode; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	//check the response body
	expected := "Hello, World!"
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}
	if string(bodyBytes) != expected {
        t.Errorf("handler returned unexpected body: got %v want %v", string(bodyBytes), expected)
    }
}
func TestGetUsersHandler(t *testing.T) {
    // ARRANGE
	// Create a new HTTP request
    req, err := http.NewRequest("GET", "/api/users", nil)
    if err != nil {
        t.Fatal(err)
    }

    // Create a new response recorder
    rr := httptest.NewRecorder()

    // Create a handler
    handler := http.HandlerFunc(handleUsers)

	//Arrange our expected response
	expected := []db.User{
			{ID: 1, Name: "User 1"},
			{ID: 2, Name: "User 2"},
			{ID: 3, Name: "User 3"},
	}
	
	expectedJSON, err := json.Marshal(expected)
	if err != nil {
		t.Fatalf("Failed to marshal expected JSON: %v", err)
	}


		//ACT
	handler.ServeHTTP(rr,req)

	    //ASSERT
    // Check the status code
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    // Check the response body
	//use of reflect lib for comparison

    if rr.Body.String() != string(expectedJSON) {
        t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
    }
	//hmm this checks actualy json instead of a stringified version
	var actual []db.User
    if err := json.Unmarshal(rr.Body.Bytes(), &actual); err != nil {
        t.Fatalf("Failed to unmarshal response body: %v", err)
    }

    if !reflect.DeepEqual(actual, expected) {
        t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
    }
}
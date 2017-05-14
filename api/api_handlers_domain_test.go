package api

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

var createDomainScenario = []struct {
	method             string
	path               string
	content            io.Reader
	expectedHTTPStatus int
}{
	{
		"POST",
		"/api/v0/domains",
		nil,
		http.StatusBadRequest,
	},
	{
		"POST",
		"/api/v0/domains",
		bytes.NewBuffer([]byte(`{"uid": "wsv.demo"}`)),
		http.StatusBadRequest,
	},
}

func TestCreateDomainHandler(t *testing.T) {
	for _, scenario := range createDomainScenario {
		log.Print("Running scenario for ", scenario)
		req, err := http.NewRequest(scenario.method, scenario.path, scenario.content)
		if err != nil {
			log.Fatal("Something failed while running the test.")
			t.Fatal(err)
		}
		recorder := httptest.NewRecorder()
		handler := http.HandlerFunc(createDomainHandler)
		handler.ServeHTTP(recorder, req)
		if status := recorder.Code; status != scenario.expectedHTTPStatus {
			t.Errorf("handler returned wrong status code: got %v want %v", status, scenario.expectedHTTPStatus)
		}
	}

}

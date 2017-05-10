package api

import (

	"net/http"
	"testing"
)

//TestDefaultHandler - TestDefaultHandler
func TestDefaultHandler(t *testing.T) {

	var w http.ResponseWriter = nil
	var r *http.Request = nil

	DefaultHandler(w,r)
/**

	req, err := http.NewRequest("GET", "/api/v0/vehicles", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DefaultHandler)



	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
**/

}

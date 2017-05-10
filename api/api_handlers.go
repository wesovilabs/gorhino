package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

//Todo -
type Todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Created   time.Time `json:"createdon"`
}

//DefaultHandler -
func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	log.Printf("%s\t%s\t%s\t%s",
		r.RemoteAddr,
		r.Method,
		r.RequestURI,
		r.Proto,
	)
	var todo Todo
	todo.Created = time.Now()
	response, err := json.MarshalIndent([]Todo{todo}, "", "    ")
	if err != nil {

	}
	w.Write(response)
}

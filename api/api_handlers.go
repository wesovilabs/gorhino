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


type jsonInterface interface {
	marshal(v interface{}, prefix, indent string) ([]byte, error)
}

type jsonMock struct {
		result	string
}

type JsonMarshal struct {

}

func (JsonMarshal) marshal(v interface{}, prefix, indent string) ([]byte, error){
	return json.MarshalIndent(v,prefix,indent)
}

func (*jsonMock) marshal(v interface{}, prefix, indent string) ([]byte, error){
	return []byte("{\"demo\":\"a\"}"),nil
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
	response, err := JsonMarshal{}.marshal([]Todo{todo}, "", "    ")
	if err != nil {

	}
	w.Write(response)
}

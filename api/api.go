package api

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"log"
	"net/http"
)
//Measurement object define
type Measurement struct {
	Measurement       int `json:"measurement"`

}
//NewMeasurement new object
func NewMeasurement()*Measurement{
	 m:= Measurement{
		Measurement:        1,
	}
	return &m
}

//FrontendAPI xxxxd
func FrontendAPI(m *Measurement) func(router chi.Router) {
	return func(router chi.Router) {
		router.Get("/get", getMessagesHandler(m))
		router.Post("/post", postMessageHandler(m))
	}
}

//GetMessages API  messages get provider
func getMessagesHandler(m *Measurement) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		response, err := json.Marshal(m)
		if err != nil {
			log.Println("Can't marshal data: ", err)
		}

		w.Write(response)

		w.WriteHeader(http.StatusOK)
	}
}
//PostMessage posts
func postMessageHandler(m *Measurement) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		measurement := NewMeasurement()
		json.NewDecoder(r.Body).Decode(measurement)

		m.Measurement = measurement.Measurement
		
		w.WriteHeader(http.StatusCreated)
		fmt.Println(m)
	}
}
//NotFound 404
func NotFound(w http.ResponseWriter, r *http.Request) {
fmt.Println("404")
	w.WriteHeader(http.StatusNotFound)
}
//CheckHealth check
func  CheckHealth(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusOK)

}
package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gobuffalo/packr"
	"github.com/gorilla/mux"
)

// Data is the container for some example data
type Data struct {
	Value string `json:"value,omitempty"`
}

// AddRequest contains the request data for the add command
type AddRequest struct {
	A int `json:"a,omitempty"`
	B int `json:"b,omitempty"`
}

// AddResponse contains the response data for the add command
type AddResponse struct {
	Result int `json:"result,omitempty"`
}

var data []Data

func getData(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func add(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t AddRequest
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	var res AddResponse
	res.Result = t.A + t.B
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func apiRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/data", getData).Methods("GET")
	router.HandleFunc("/api/add", add).Methods("POST")
	return router
}

func main() {

	data = append(data, Data{Value: "Hi there!"})

	router := apiRouter()
	http.Handle("/api/", router)

	box := packr.NewBox("./client")
	fs := http.FileServer(box)
	http.Handle("/", fs)
	log.Println("Listening on port 3001 ...")
	log.Fatal(http.ListenAndServe(":3001", nil))
}

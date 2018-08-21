package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Data struct {
	Value string `json:"value,omitempty"`
}

var data []Data

func getData(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(data)
}

func apiRouter() *mux.Router {
	router := mux.NewRouter()
	router = mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/data", getData).Methods("GET")
	return router
}

func main() {

	data = append(data, Data{Value: "Hi"})

	// router := mux.NewRouter()
	router := apiRouter()
	http.Handle("/api/", router)

	fs := http.FileServer(http.Dir("client"))
	http.Handle("/", fs)
	// http.ListenAndServe(":3001", nil)
	log.Println("Listening on port 8000 ...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

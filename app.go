package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gobuffalo/packr"
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

	router := apiRouter()
	http.Handle("/api/", router)

	box := packr.NewBox("./client")
	fs := http.FileServer(box)
	http.Handle("/", fs)
	log.Println("Listening on port 3001 ...")
	log.Fatal(http.ListenAndServe(":3001", nil))
}

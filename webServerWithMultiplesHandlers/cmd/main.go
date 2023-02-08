package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	//myweb/model"
)

// vars
var result []string

// vars para lo que sería el package model
type CreateVolume struct {
	Zone        string `json:"zone"`
	Application string `json:"application"`
	Environment string `json:"environment"`
	Component   string `json:"component"`
	Pv          string `json:"pv"`
}

type Response struct {
	Response string `json:"response"`
	Code     int    `json:"code"`
}

// end vars of package model

// Web development
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello!")
}

// post handler
func createHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/create" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	var vol CreateVolume
	err := json.NewDecoder(r.Body).Decode(&vol)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		fmt.Printf("Estoy dentro del else.\n")
		fmt.Printf("%s\n", vol)
	}

	fmt.Fprint(w, "POST done")
}

// function del package model
func CreateVol(body CreateVolume) ([]byte, error) {
	fmt.Printf("Entro en la funcion de crear\n")
	res_json := &Response{"OK", 200}
	res, _ := json.Marshal((res_json))
	fmt.Printf("Salgo de la funcion de crear\n")
	return res, nil
}

// end package model

// Main
func main() {
	// web init
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/create", createHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

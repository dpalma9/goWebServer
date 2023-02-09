package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
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
		validateParam, paramErr := CreateIsValid(vol)
		if paramErr != nil {
			http.Error(w, string(validateParam), http.StatusBadRequest)
			//fmt.Fprint(w, string(validateParam))
			return
		}
		response, _ := CreateVol(vol)
		fmt.Printf("Respuesta de creacion.\n")
		fmt.Printf("%s\n", response)
		fmt.Fprint(w, string(response))
	}

	fmt.Fprint(w, "POST done")
}

// function del package model
func CreateVol(body CreateVolume) ([]byte, error) {
	fmt.Printf("Entro en la funcion de crear\n")
	root_path := "/tmp"
	pv_path := root_path + "/" + body.Application + "/" + body.Environment + "/" + body.Component + "-" + body.Pv
	fmt.Printf("Vamos a crear el siguiente path --> %s\n", pv_path)
	if err := os.MkdirAll(pv_path, os.ModePerm); err != nil {
		log.Fatal(err)
	}
	res_json := &Response{"OK", 200}
	res, _ := json.Marshal((res_json))
	fmt.Printf("Salgo de la funcion de crear\n")
	return res, nil
}

func CreateIsValid(body CreateVolume) ([]byte, error) {
	fmt.Printf("Entro en la funcion de validar\n")
	// response vars
	res_ok := &Response{"Parameters are OK", 200}
	res, _ := json.Marshal((res_ok))
	res_missing := &Response{"Some parameter is missing", 409}
	res_miss, _ := json.Marshal((res_missing))

	// valid values vars
	validZone := map[string]bool{
		"intranet": true,
		"internet": true,
		"platform": true,
	}

	validEnv := map[string]bool{
		"pre":  true,
		"pro":  true,
		"tst":  true,
		"dev":  true,
		"sbox": true,
	}

	// Check if all parameters are on the request
	if body.Application == "" {
		return res_miss, fmt.Errorf("'application' is missing")
	}
	if body.Zone == "" {
		return res_miss, fmt.Errorf("'zone' is missing")
	}
	if body.Environment == "" {
		return res_miss, fmt.Errorf("'environment' is missing")
	}
	if body.Component == "" {
		return res_miss, fmt.Errorf("'component' is missing")
	}
	if body.Pv == "" {
		return res_miss, fmt.Errorf("'pv' is missing")
	}

	// Check if values are valid
	if !validZone[strings.ToLower(body.Zone)] {
		return res_miss, fmt.Errorf("'zone' hasn't a valid value")
	}
	if !validEnv[strings.ToLower(body.Environment)] {
		return res_miss, fmt.Errorf("'environment' hasn't a valid value")
	}

	fmt.Printf("Salgo de la funcion de validar\n")
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

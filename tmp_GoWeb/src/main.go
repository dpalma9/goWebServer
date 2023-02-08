package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//vars:
type Response struct {
	Response string `json:"response"`
	Code     int    `json:"code"`
}

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

	response, err := getJsonResponse()
	if err != nil {
		panic(err)
	}
	//res, _ := json.Marshal(response)
	//fmt.Fprintf(w, string(response))
	//fmt.Fprintf(w, "Hello!")
	//fmt.Printf("%s", res)
	fmt.Printf("%s\n", response)
	fmt.Printf("Termino peticion.\n")
}

func main() {
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func getJsonResponse() ([]byte, error) {
	fmt.Printf("Entro en get response\n")

	respuesta := &Response{"hola", 200}
	res, err := json.Marshal(respuesta)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(string(res))
	fmt.Printf("%s\n", res)
	fmt.Printf("Salgo de get response\n")

	return (res), err
}

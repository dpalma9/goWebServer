package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"myweb/model"
)

// vars
var result []string

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

	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)
		} else {
			result = model.CreateVol(body)
			fmt.Printf("Print de result\n")
			fmt.Printf(result)
		}
		//results = append(results, string(body))

		fmt.Fprint(w, "POST done")

	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}

}

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

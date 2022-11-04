package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"github.com/robfig/cron"
)

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

func autoConsume() {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}

	// appending to existing query args
	q := req.URL.Query()

	// assign encoded query string to http request
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return
	}

	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.Status)
	fmt.Println(string(responseBody))
}

func main() {
	// create a WaitGroup needed to control goroutine
	wg := new(sync.WaitGroup)

	// add one goroutines to `wg` WaitGroup
	wg.Add(1)

	// Web server
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080 with goroutine AND cron\n")
	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal(err)
		}
	}()

	// cron goroutine
	c := cron.New()
    //c.AddFunc("* * * * * *", autoConsume)
	tiempo := "@every 1m"
	c.AddFunc(tiempo, autoConsume)
    go c.Start()

	// wait until WaitGroup is done
	wg.Wait()
	fmt.Printf("Hola, esto es un test.\n")
}
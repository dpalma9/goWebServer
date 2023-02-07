package main

import (
    "fmt"
    "net/http"
    "log"
    "encoding/json"
)

//vars:
type Data struct {
    response string
    code int
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

    response, err := getJsonResponse();
    if err != nil {
        panic(err)
    }
    fmt.Fprintf(w, string(response))
    fmt.Fprintf(w, "Hello!")
}

func main() {
    http.HandleFunc("/hello", helloHandler)

    fmt.Printf("Starting server at port 8080\n")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}

func getJsonResponse()([]byte, error) {
    //p := Payload{d}
    fmt.Printf("Entro en get response\n")
    //respuesta := &Data{response: "hola", code: 200}
    respuesta := Data{response: "hola", code: 200}
    res, err := json.Marshal(respuesta)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(res))

    //return json.MarshalIndent(p, "", "  ")
    return json.MarshalIndent(respuesta, "", "  ")
}

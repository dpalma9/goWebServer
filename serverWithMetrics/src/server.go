package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
    "math/rand"
    //"time"
)
// --PROMETHEUS DEV--
//Define a struct for you collector that contains the pointer
//to prometheus descriptors for my metric I wish to expose.
type myCollector struct {
	myMetric *prometheus.Desc
}

//Create a constructor for collector that
//initializes every descriptor and returns a pointer to the collector
func myCollectorConstructor() *myCollector {
	return &myCollector{
		myMetric: prometheus.NewDesc("my_metric",
			"An invented metric to expose",
			nil, nil,
		),
	}
}

//For each and every collector must implement the Describe function.
//It essentially writes all descriptors to the prometheus desc channel.
func (collector *myCollector) Describe(ch chan<- *prometheus.Desc) {

	//Update this section with the each metric you create for a given collector
	ch <- collector.myMetric
}

//Collect implements required collect function for all promehteus collectors
func (collector *myCollector) Collect(ch chan<- prometheus.Metric) {

	//Implement logic here to determine proper metric value to return to prometheus
	//for each descriptor or call other functions that do so.
	var metricValue float64
        //temporal value
        metricValue += rand.Float64()

	//Write latest value for each metric in the prometheus metric channel.
	//Can pass CounterValue, GaugeValue, or UntypedValue types here.
	m1 := prometheus.MustNewConstMetric(collector.myMetric, prometheus.GaugeValue, metricValue)
	//m1 = prometheus.NewMetricWithTimestamp(time.Now().Add(-time.Hour), m1)
	ch <- m1
}
// --PROMETHEUS DEV END--
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

func main() {
    // Init prometheus collector
    promColl := myCollectorConstructor()
    prometheus.MustRegister(promColl)
    http.Handle("/metrics", promhttp.Handler())

    // web init
    fileServer := http.FileServer(http.Dir("./static"))
    http.Handle("/", fileServer)
    http.HandleFunc("/hello", helloHandler)

    fmt.Printf("Starting server at port 8080\n")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}

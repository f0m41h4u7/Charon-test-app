package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Data struct {
	Values []float64 `json:"values"`
}

var datasets []Data

var (
	testMetric = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "testMetrics",
		Help: "Some test metrics",
	})
)

func exposeMetrics() {
	for i := 1; i < 4; i++ {
		var d Data
		dat, err := ioutil.ReadFile("datasets/dataset" + strconv.Itoa(i) + ".json")
		if err != nil {
			log.Fatal(fmt.Errorf("Failed to read metrics dataset %d: %w\n", i, err))
		}
		err = json.Unmarshal([]byte(dat), &d)
		if err != nil {
			log.Fatal(fmt.Errorf("Failed to parse metrics dataset %d: %w\n", i, err))
		}
		datasets = append(datasets, d)
	}
}

func main() {
	exposeMetrics()
	for _, d := range datasets {
		go func(d Data) {
			for _, metric := range d.Values {
				testMetric.Set(metric)
				log.Print("Set metric ", metric)
				time.Sleep(2 * time.Minute)
			}
		}(d)
	}

	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":1337", nil)
	if err != nil {
		log.Fatal(fmt.Errorf("Server error: %w\n", err))
	}
}

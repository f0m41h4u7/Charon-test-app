package main

import (
        "net/http"
        "time"
        "log"
        "io/ioutil"
        "encoding/json"
        "github.com/prometheus/client_golang/prometheus"
        "github.com/prometheus/client_golang/prometheus/promauto"
        "github.com/prometheus/client_golang/prometheus/promhttp"
)

type Data struct {
        Values  []float64       `json:"values"`
}

var dataset Data

func recordMetrics() {
        go func() {
                for _, metric := range dataset.Values{
                        testMetric.Set(metric)
			log.Print("Set metric ", metric)
                        time.Sleep(2 * time.Minute)
                }
        }()
}

var (
        testMetric = promauto.NewGauge(prometheus.GaugeOpts{
                Name: "testMetrics",
                Help: "Some test metrics",
        })
)

func main() {
        dat, err := ioutil.ReadFile("dataset.json")
        if err != nil {
                log.Fatal(err)
        }
        err = json.Unmarshal([]byte(dat), &dataset)
        if err != nil {
                log.Fatal(err)
        }

        recordMetrics()

        http.Handle("/metrics", promhttp.Handler())
        http.ListenAndServe(":1337", nil)
}

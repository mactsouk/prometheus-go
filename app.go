package main

import (
	"net/http"

	"log"
	"math/rand"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// PORT is the TCP port number the server will listen to
var PORT = ":1234"

var (
	counter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "golang",
			Name:      "myCounter",
			Help:      "This is my counter",
		})

	gauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "golang",
			Name:      "myGauge",
			Help:      "This is my gauge",
		})

	histogram = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: "golang",
			Name:      "myHistogram",
			Help:      "This is my histogram",
		})

	summary = prometheus.NewSummary(
		prometheus.SummaryOpts{
			Namespace: "golang",
			Name:      "mySummary",
			Help:      "This is my summary",
		})
)

func main() {

	rand.Seed(time.Now().Unix())

	http.Handle("/metrics", promhttp.Handler())

	prometheus.MustRegister(counter)
	prometheus.MustRegister(gauge)
	prometheus.MustRegister(histogram)
	prometheus.MustRegister(summary)

	go func() {
		for {
			counter.Add(rand.Float64() * 5)
			gauge.Add(rand.Float64()*15 - 5)
			histogram.Observe(rand.Float64() * 10)
			summary.Observe(rand.Float64() * 10)

			time.Sleep(time.Second)
		}
	}()

	log.Println("Listening to port", PORT)
	log.Println(http.ListenAndServe(PORT, nil))
}

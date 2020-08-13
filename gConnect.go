package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// PORT is the TCP port number the server will listen to
var PORT = ":2345"

var (
	counter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "custom",
			Name:      "my_counter",
			Help:      "# of emails in the Gmail account",
		})
)

func main() {
	rand.Seed(time.Now().Unix())
	http.Handle("/metrics", promhttp.Handler())
	prometheus.MustRegister(counter)

	go func() {
		for {
			counter.Add(rand.Float64() * 5)

			time.Sleep(5 * time.Second)
		}
	}()

	log.Println("Listening to port", PORT)
	log.Println(http.ListenAndServe(PORT, nil))
}

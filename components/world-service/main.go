package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var addr = flag.String("listen-address", ":20081", "The address to listen on for HTTP requests.")

var (
	helloCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "world_counter",
			Help: "counter for worlds",
		},
	)
)

func main() {
	flag.Parse()
	prometheus.MustRegister(helloCounter)
	go func() {
		for {
			helloCounter.Inc()
			time.Sleep(1 * time.Second)
		}
	}()
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*addr, nil))
}

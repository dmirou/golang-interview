package main

import (
	"context"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"golang.org/x/sync/errgroup"
)

var (
	signupCount = promauto.NewCounter(prometheus.CounterOpts{
		Name: "signupcounter_signup_count",
		Help: "The total number of signups",
	})

	onlineUsers = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "signupcounter_online_users",
		Help: "Count of online users",
	})

	runningJobs = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "signupcounter_running_jobs",
			Help: "Count of running jobs",
		},
		[]string{"type"},
	)
)

func main() {

	rand.Seed(time.Now().UnixNano())

	g, _ := errgroup.WithContext(context.Background())

	go func() {
		for {
			select {
			case <-time.After(5 * time.Second):
				onlineUsers.Set(float64(20 + rand.Intn(60)))
			}
		}
	}()

	go func() {
		for {
			select {
			case <-time.After(5 * time.Second):
				runningJobs.WithLabelValues("notification").Set(float64(20 + rand.Intn(60)))
				runningJobs.WithLabelValues("completeOrder").Set(float64(rand.Intn(60)))
			}
		}
	}()

	infoMux := http.NewServeMux()
	infoMux.Handle("/metrics", promhttp.Handler())
	infoSrv := http.Server{
		Addr:    ":2112",
		Handler: infoMux,
	}

	g.Go(func() error {
		if err := infoSrv.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				log.Printf("info server unexpected error: %v", err)
			}
		}

		return nil
	})

	apiMux := http.NewServeMux()
	apiMux.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		signupCount.Inc()
		io.WriteString(w, "signup registered")
	})
	apiSrv := http.Server{
		Addr:    ":8080",
		Handler: apiMux,
	}

	g.Go(func() error {
		if err := apiSrv.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				log.Printf("api server unexpected error: %v", err)
			}
		}

		return nil
	})

	g.Wait()

	// curl localhost:8080/signup
	// curl localhost:2112/metrics
}

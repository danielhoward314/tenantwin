package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		userSvcRes, err := http.Get("http://localhost:9090/api/v1/health")
		if err != nil {
			log.Printf("userService err\n%v\n", err)
		}
		paymentSvcRes, err := http.Get("http://localhost:9091/api/v1/health")
		if err != nil {
			log.Printf("paymentService err\n%v\n", err)
		}
		leaseSvcRes, err := http.Get("http://localhost:9092/api/v1/health")
		if err != nil {
			log.Printf("leaseService err\n%v\n", err)
		}
		requestSvcRes, err := http.Get("http://localhost:9093/api/v1/health")
		if err != nil {
			log.Printf("requestService err\n%v\n", err)
		}
		mr := io.MultiReader(userSvcRes.Body, paymentSvcRes.Body, leaseSvcRes.Body, requestSvcRes.Body)
		io.Copy(w, mr)
	}).Methods(http.MethodGet)
	allowList := []string{
		"http://localhost:9090",
		"http://localhost:9091",
		"http://localhost:9092",
		"http://localhost:9093",
	}
	ch := gohandlers.CORS(gohandlers.AllowedOrigins(allowList))
	s := &http.Server{
		Addr:         ":8080",
		Handler:      ch(r),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.Fatal("API Composer unable to start server")
		}
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	sig := <-c
	log.Printf("Got an os.Signal: %v", sig)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	s.Shutdown(ctx)
}

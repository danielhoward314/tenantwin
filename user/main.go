package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"

	"github.com/danielhoward314/tenantwin/user/handler"
	pg "github.com/danielhoward314/tenantwin/user/repository"
	"github.com/danielhoward314/tenantwin/user/svc"
)

func main() {
	httpPort := os.Getenv("HTTP_PORT")
	postgresHost := os.Getenv("POSTGRES_HOST")
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresDBName := os.Getenv("POSTGRES_DB_NAME")
	postgresPort := os.Getenv("POSTGRES_PORT")
	repo, err := pg.NewUserRepository(postgresHost, postgresUser, postgresPassword, postgresDBName, postgresPort)
	if err != nil {
		log.Fatal(err)
	}
	us := svc.NewUserService(repo)
	l := log.New(os.Stdout, "userService", log.LstdFlags)
	uh := handler.NewUserHandler(l, us)
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	users := api.PathPrefix("/users").Subrouter()
	users.HandleFunc("/signup", uh.Signup).Methods(http.MethodPost)
	allowList := []string{
		"http://localhost:8080",
		"http://localhost:9091",
		"http://localhost:9092",
		"http://localhost:9093",
		"http://stg-tenantwin.com:8080",
		"http://stg-tenantwin.com:9091",
		"http://stg-tenantwin.com:9092",
		"http://stg-tenantwin.com:9093",
		"http://tenantwin.com:8080",
		"http://tenantwin.com:9091",
		"http://tenantwin.com:9092",
		"http://tenantwin.com:9093",
	}
	ch := gohandlers.CORS(gohandlers.AllowedOrigins(allowList))
	s := &http.Server{
		Addr:         fmt.Sprintf(":%v", httpPort),
		Handler:      ch(r),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.Fatal("userService unable to start server")
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

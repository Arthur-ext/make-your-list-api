package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
)

type server struct {
	http.Server
}

func NewServer() server {
	if os.Getenv("APPENV") == "" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("some error occured. Err: %s", err)
		}
	}
	port := os.Getenv("PORT")
	if port != "" {
		port = os.Getenv("LOCAL_PORT")
	}

	log.Println("configuring server...")

	api := newAPI()

	srv := http.Server{
		Addr: port,
		Handler: api.router,
	}

	return server{srv}
}

func (s server) Start() {
	log.Println("starting server...")

	go func() {
		if err := s.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()

	log.Printf("Listening on %s\n", s.Addr)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	sig := <-quit

	log.Println("shutting down server... Reason: ", sig)

	if err := s.Shutdown(context.Background()); err != nil {
		panic(err)
	}
	log.Println("server gracefully stopped")
}
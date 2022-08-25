package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
)

type server struct {
	http.Server
}

func NewServer() server {
	var port string

	if os.Getenv("APPENV") == "prod" {
		port = os.Getenv("PORT")
	} else {
		port = os.Getenv("LOCAL_PORT")
	}

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	log.Println("configuring server...")

	api := newAPI()

	srv := http.Server{
		Addr: ":"+port,
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
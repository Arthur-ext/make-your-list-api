package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"wedding_gifts/api/router"
)

type server struct {
	http.Server
}

func NewServer() server {
	log.Println("configuring server...")

	api := router.InitRouter()

	srv := http.Server{
		Addr: "9000",
		Handler: api,
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
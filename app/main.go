package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/julienschmidt/httprouter"
)

func newRouter() *httprouter.Router {
	mux := httprouter.New()

	ytApiKey := os.Getenv("YOUTUBE_API_KEY")
	if ytApiKey == "" {
		log.Fatal("Api key not found")
	}

	mux.GET("/youtube.com/channel/stats", getChannelStats())

	return mux
}

func main() {
	srv := &http.Server{
		Addr:    ":1010",
		Handler: newRouter(),
	}
	idleConnectionClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		signal.Notify(sigint, syscall.SIGTERM)
		<-sigint

		log.Println("service interrupt recieved")

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)

		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			log.Fatalf("server shutdown failed %v", err)
		}

		log.Println("Shutdown closed successfully")

		close(idleConnectionClosed)

	}()

	if err := srv.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("fatal http server failed to start %v", err)
		}
	}
	fmt.Println("Server has started")
	<-idleConnectionClosed
	log.Println("Service Stop")

}
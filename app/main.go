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

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Home Page")
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", home)

	myServer := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	idleConnectionClosed := make(chan struct{})
	go func() {
		interruptSignal := make(chan os.Signal, 1)

		signal.Notify(interruptSignal, os.Interrupt)
		signal.Notify(interruptSignal, syscall.SIGTERM)
		<-interruptSignal

		log.Println("Service Interruppt Received")

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		if err := myServer.Shutdown(ctx); err != nil {
			log.Println("Error While Shutting Down Server", err)
		}

		log.Println("Shutdown Done Succesfuly")

		close(idleConnectionClosed)
	}()

	log.Println("Starting Service at Port 8080")
	if err := myServer.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatalln("Server Failed to Start", err)
		}
	}

	<-idleConnectionClosed
	log.Println("Service Stopped Succesfully")
}

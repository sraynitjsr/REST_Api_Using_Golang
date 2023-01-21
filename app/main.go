package main

import (
	"fmt"
	"net/http"

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

	myServer.ListenAndServe()
}

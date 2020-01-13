package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome 1!")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	http.ListenAndServe(":8080", router)
}

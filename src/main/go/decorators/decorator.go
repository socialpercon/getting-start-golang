package main

import (
	"fmt"
	"net/http"
)

func authRequried(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.FormValue("user") == "" {
			http.Error(w, "unknown user", http.StatusForbidden)
			return
		}
		f(w, r)
	}
}

func main() {
	var hiHandler = authRequried(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hi, %v", r.FormValue("user"))
		},
	)
	http.HandleFunc("/hi", hiHandler)
	http.ListenAndServe(":8080", nil)
}

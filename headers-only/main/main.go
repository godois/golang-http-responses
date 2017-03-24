package main

import (
"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":3000", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", "Web Server in golang")
	w.WriteHeader(200)
}
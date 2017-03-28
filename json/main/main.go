package main

import (
	"encoding/json"
	"net/http"
)

type Person struct {
	Name    string
	Hobbies []string
}

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":3000", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	profile := Person{"Marcio Godoi", []string{"programming", "gamming"}}

	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
package main

import (
	"encoding/xml"
	"net/http"
)

type Profile struct {
	Name    string
	Hobbies []string `xml:"Hobbies>Hobby"`
}

func main() {
	http.HandleFunc("/", responseBuffer)
	http.ListenAndServe(":3000", nil)
}

func responseBuffer(w http.ResponseWriter, r *http.Request) {

	profile := Profile{"Marcio Godoi", []string{"cooking", "programming"}}

	x, err := xml.MarshalIndent(profile, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/xml")
	w.Write(x)
}

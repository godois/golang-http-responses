package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", responseBuffer)
	http.ListenAndServe(":3000", nil)
}

func responseBuffer(w http.ResponseWriter, r *http.Request) {
	// Assuming you want to serve a photo at 'images/foo.png'
	fp := "mushroom.png"
	http.ServeFile(w, r, fp)
}

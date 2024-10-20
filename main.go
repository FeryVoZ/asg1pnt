package main

import (
	"io"
	"net/http"
)

func endpoint1(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World! This is endpoint 1...")
}

func endpoint2(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World! This is endpoint 2...")
}

func getroot(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Choose between /ep1 or /ep2")
}

func main() {

	http.HandleFunc("/", getroot)
	http.HandleFunc("/ep1", endpoint1)
	http.HandleFunc("/ep2", endpoint2)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}

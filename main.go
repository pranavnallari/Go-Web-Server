package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err :%v\n", err)
		return
	}
	fmt.Fprintf(w, "POST request Successful\n")
	s := r.FormValue("s")
	fmt.Fprintf(w, "S = %s\n", s)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method used is incorrect !", http.StatusNotFound)
		return
	}
}

func main() {
	fileServer := http.FileServer(http.Dir("./static")) //Initializes server serving over the static directory
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server at port 3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}

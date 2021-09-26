package main

import (
"fmt"
"log"
"net/http"
"os"
)

func hostname() string {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	return name
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	hostname := hostname()

	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	fmt.Printf("You machine is : %s!\n", hostname)
	fmt.Fprintf(w, "Hello World ! You machine is : %s!", hostname)
}

func main() {
	http.HandleFunc("/hello", helloHandler) // Update this line of code
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

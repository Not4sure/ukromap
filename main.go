package main

import (
	"fmt"
	"log"
	"net/http"
)

func test(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hello")
}

func main() {
	http.HandleFunc("/", test)
	fmt.Println("Hello")
	log.Fatal(http.ListenAndServe(":80", nil))
}
package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Printf("running Main()")
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)
	//err := http.ListenAndServe(":3333", nil)
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "this is my website!\n")
}
func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

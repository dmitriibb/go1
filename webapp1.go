package main

import (
	"fmt"
	"log"
	"net/http"
	"web-app1/controller"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/hello", handlerHello)
	http.HandleFunc("/person", controller.HandlerUser)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

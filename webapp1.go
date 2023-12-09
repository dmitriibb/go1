package main

import (
	"fmt"
	"log"
	"net/http"
	"web-app1/common/db"
	"web-app1/user"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func main() {
	db.TestConnectPostgres()
	http.HandleFunc("/", handler)
	http.HandleFunc("/hello", handlerHello)
	http.HandleFunc("/user", user.HandlerUser)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

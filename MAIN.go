package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", sayhelloName) // setting router rule

	http.HandleFunc("/login", login)
	http.HandleFunc("/upload", upload)

	err := http.ListenAndServe(":9090", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

//http://127.0.0.1:9090/login
//http://127.0.0.1:9090/upload
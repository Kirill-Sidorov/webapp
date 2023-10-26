package main

import (
	"log"
	"net/http"
	"webapp/handlers/mainpage"
	"webapp/handlers/session"
)

func main2() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", mainpage.MainPageHandler)
	mux.HandleFunc("/main", mainpage.MainPageHandler)
	mux.HandleFunc("/login", session.LoginHandler)

	log.Println("Start server")
	err := http.ListenAndServe("localhost:8080", mux)
	log.Fatal(err)
}
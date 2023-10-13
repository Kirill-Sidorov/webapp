package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/handlers/mainpage"
	"webapp/handlers/login"
)

func main() {
	fmt.Print("Start")

	http.HandleFunc("/main", mainpage.MainPageHandler)
	http.HandleFunc("/login", login.LoginHandler)

	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
package session

import (
	"fmt"
	"log"
	"net/http"
)

func LoginHandler(writer http.ResponseWriter, request *http.Request) {
	_, err := writer.Write([]byte("Login page handler"))

	if err != nil {
		log.Print(err)
	}
}

func DestroySessions() {
	fmt.Println("Sessions destoryed!")
}
package mainpage

import (
	"log"
	"net/http"
)

func MainPageHandler(writer http.ResponseWriter, request *http.Request) {
	_, err := writer.Write([]byte("MainPageHandler"))

	if err != nil {
		log.Print(err)
	}
}
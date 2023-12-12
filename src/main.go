package main

import (
	"digitalwallet-service/src/webapp/config"
	"digitalwallet-service/src/webapp/route"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Setup()
	router := route.CreateRouter()

	port := fmt.Sprintf(":%d", config.ApiConfig.Port)

	println("Listening ", port)
	log.Fatal(http.ListenAndServe(port, router))
}

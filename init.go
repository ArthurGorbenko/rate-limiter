package main

import (
	"fmt"
	"net/http"
)

var defaultTransport = http.DefaultTransport

const targetProxy = "http://localhost:3000"

func main() {
	setInterval(handleTokenRefil, REFIL_INTERVAL)

	server := http.Server{
		Addr:    ":8080",
		Handler: rateLimiter(http.HandlerFunc(HandleProxy)),
	}

	fmt.Println("Starting the server!")
	err := server.ListenAndServe()

	if err != nil {
		fmt.Println("Error starting the server: ", err)
	}
}

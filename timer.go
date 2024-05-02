package main

import (
	"net/http"
	"time"
)

const REFIL_INTERVAL = 2 * time.Second

var tokenCount = 4

func handleTokenRefil() {
	if tokenCount == 4 {
		return
	}

	tokenCount += 2

	if tokenCount > 4 {
		tokenCount = 4
	}

	println("Token count: ", tokenCount)
}

func rateLimiter(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if tokenCount == 0 {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			tokenCount = 4
			return
		}

		tokenCount--

		next.ServeHTTP(w, r)
	})
}

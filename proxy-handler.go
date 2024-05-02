package main

import (
	"io"
	"net/http"
)

func HandleProxy(w http.ResponseWriter, r *http.Request) {

	targetURL := targetProxy + r.URL.String()

	proxyReq, err := http.NewRequest(r.Method, targetURL, r.Body)
	if err != nil {
		println("Error creating proxy request: ", err)
		http.Error(w, "Error creating proxy request", http.StatusInternalServerError)
		return
	}

	for name, values := range r.Header {
		for _, value := range values {
			proxyReq.Header.Add(name, value)
		}
	}

	resp, err := defaultTransport.RoundTrip(proxyReq)

	if err != nil {
		println("Error fetching response from upstream server: ", err.Error())
		http.Error(w, "Error fetching response from upstream server", http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	for name, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(name, value)
		}
	}

	w.WriteHeader(resp.StatusCode)

	io.Copy(w, resp.Body)
}

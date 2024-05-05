package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
)

type EndpointStats struct {
	Count int
	mu    sync.Mutex
}

func (es *EndpointStats) Increment() {
	es.mu.Lock()
	defer es.mu.Unlock()
	es.Count++
}

func (es *EndpointStats) GetCount() int {
	es.mu.Lock()
	defer es.mu.Unlock()
	return es.Count
}

func handler(endpointName string, stats *EndpointStats) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Increment the hit count for the endpoint
		if headerValue := r.Header.Get("Programmatic-Request"); strings.ToLower(headerValue) == "true" {
			// Increment the hit count for the endpoint
			stats.Increment()
		}

		// Display the request details
		requestDetails := fmt.Sprintf("Request Method: %s\nRequest URL: %s\nRequest Headers: %v\n", r.Method, r.URL.String(), r.Header)

		// Return the hit count and request details
		response := fmt.Sprintf("Endpoint: %s\nHit Count: %d\n\n%s", endpointName, stats.GetCount(), requestDetails)

		// Write the response to the client
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
	}
}

func main() {
	// Create stats objects for each endpoint
	endpoint1Stats := &EndpointStats{}
	endpoint2Stats := &EndpointStats{}
	endpoint3Stats := &EndpointStats{}

	// Create an HTTP server with handlers for each endpoint
	http.HandleFunc("/endpoint1", handler("Endpoint1", endpoint1Stats))
	http.HandleFunc("/endpoint2", handler("Endpoint2", endpoint2Stats))
	http.HandleFunc("/endpoint3", handler("Endpoint3", endpoint3Stats))

	// Start the HTTP server on port 8090
	port := ":8090"
	fmt.Printf("Starting server on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

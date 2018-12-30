package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

func httpClient(requestTimeout time.Duration) *http.Client {
	return &http.Client{
		Timeout: requestTimeout,
	}
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "80"
	}

	clGeter := New(httpClient(1 * time.Second))

	log.Fatal(http.ListenAndServe(":"+port, router(clGeter)))
}

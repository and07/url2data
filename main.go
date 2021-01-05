package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

const requestTimeout = 5 * time.Second

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

	clGeter := New(httpClient(requestTimeout))

	log.Fatal(http.ListenAndServe(":"+port, router(clGeter)))
}

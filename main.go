package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/chromedp/chromedp"
)

const requestTimeout = 5 * time.Second

func httpClient(requestTimeout time.Duration) *http.Client {
	return &http.Client{
		Timeout: requestTimeout,
	}
}

func main() {

	// create chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	port := os.Getenv("PORT")

	if port == "" {
		port = "80"
	}

	clGeter := New(httpClient(requestTimeout))

	log.Fatal(http.ListenAndServe(":"+port, router(ctx, clGeter)))
}

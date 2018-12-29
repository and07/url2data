package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

const testURL = "127.0.0.1"

// RoundTripFunc .
type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip .
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

//NewTestClient returns *http.Client with Transport replaced to avoid making real calls
func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

//TestDoStuffWithRoundTripper
func TestDoStuffWithRoundTripper(t *testing.T) {

	client := NewTestClient(func(req *http.Request) *http.Response {

		text := "test test"

		// Test request parameters
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(text)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})

	api := New(client)
	_, errAPIGet := api.Get(testURL)
	if errAPIGet != nil {
		t.Errorf("%#v", errAPIGet)
	}

}

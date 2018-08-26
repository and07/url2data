package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/julienschmidt/httprouter"
)

func getData4Url(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err.Error()
	}
	responseString := string(responseData)
	return responseString
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Not protected!\n")
}

func urlData(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var url string = `{"error" : "require url param"}`
	var callback string = "callback"
	if len(r.URL.RawQuery) > 0 {

		queryValues := r.URL.Query()
		url = queryValues.Get("url")
		if url == "" {
			w.WriteHeader(400)
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			fmt.Fprintf(w, "%s", url)
			return
		}

		callback = queryValues.Get("jsonp")
		if callback == "" {
			callback = "callback"
		}

	} else {
		w.WriteHeader(400)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		fmt.Fprintf(w, "%s", url)
		return
	}
	w.WriteHeader(200)
	//w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(w, "%s(`{ 'res' : '%s'}`)", callback, getData4Url(url))
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router := httprouter.New()
	router.GET("/", index)
	router.GET("/data", urlData)

	log.Fatal(http.ListenAndServe(":"+port, router))
}

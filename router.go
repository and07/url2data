package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Not protected!\n")
}

func urlData(cl Geter) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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
		res, errGet := cl.Get(url)
		if errGet != nil {
			w.WriteHeader(400)
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			fmt.Fprintf(w, "errGet %s", url)
			return
		}

		w.WriteHeader(200)
		//w.Header().Set("Content-Type", "application/json; charset=utf-8")
		htmlString := strings.ReplaceAll(*res, "`", "'")
		/*
			doc, err := html.Parse(strings.NewReader(htmlString))
			if err != nil {
				log.Fatal(err)
			}

			buf := bytes.NewBuffer([]byte{})
			if err := html.Render(buf, doc); err != nil {
				log.Fatal(err)
			}
			log.Println(buf.String())
		*/
		htmlString = removeScriptsLansana(htmlString)
		fmt.Fprintf(w, "%s(`{ 'res' : '%s'}`)", callback, htmlString)
	}
}

func router(cl Geter) *httprouter.Router {
	router := httprouter.New()
	router.GET("/", index)
	router.GET("/data", urlData(cl))
	return router
}

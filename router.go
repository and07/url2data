package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/chromedp"
	"github.com/julienschmidt/httprouter"
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Not protected!\n")
}

func urlContent(ctx context.Context, cl Geter) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

		var res string

		err := chromedp.Run(ctx,
			emulation.SetUserAgentOverride("WebScraper 1.0"),
			chromedp.Navigate(url),
			chromedp.Sleep(2*time.Second),
			chromedp.ActionFunc(func(ctx context.Context) error {
				node, err := dom.GetDocument().Do(ctx)
				if err != nil {
					return err
				}
				res, err = dom.GetOuterHTML().WithNodeID(node.NodeID).Do(ctx)
				return err
			}),
		)

		if err != nil {
			log.Fatal(err)
			w.WriteHeader(400)
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			fmt.Fprintf(w, "errGet %s", url)
			return
		}

		w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
		w.WriteHeader(200)
		var htmlString string
		//htmlString = res

		htmlString = strings.ReplaceAll(res, "\"", "'")
		htmlString = strings.ReplaceAll(htmlString, "\r\n", "")
		htmlString = strings.ReplaceAll(htmlString, "\t", "")
		htmlString = strings.ReplaceAll(htmlString, "\n", "")

		htmlString = removeScriptsLansana(htmlString)

		fmt.Fprintf(w, `%s({"res":"%s","url":"%s"})`, callback, htmlString, url)
		//fmt.Fprintf(w, htmlString)
	}
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

		w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
		w.WriteHeader(200)
		htmlString := strings.ReplaceAll(*res, "\"", "'")
		htmlString = strings.ReplaceAll(htmlString, "\r\n", "")
		htmlString = strings.ReplaceAll(htmlString, "\t", "")
		htmlString = strings.ReplaceAll(htmlString, "\n", "")
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
		fmt.Fprintf(w, `%s({"res":"%s","url":"%s"})`, callback, htmlString, url)
	}
}

func router(ctx context.Context, cl Geter) *httprouter.Router {
	router := httprouter.New()
	router.GET("/", index)
	router.GET("/data", urlData(cl))
	router.GET("/content", urlContent(ctx, cl))
	return router
}

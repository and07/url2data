package main

import(
	"golang.org/x/net/html"
)

func removeScript(n *html.Node) {
    // if note is script tag
    if n.Type == html.ElementNode && n.Data == "script" {
        n.Parent.RemoveChild(n)
        return // script tag is gone...
    }
    // traverse DOM
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        removeScript(c)
    }
}

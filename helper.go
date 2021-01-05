package main

import (
	"strings"

	"golang.org/x/net/html"
)

func removeScriptsLansana(s string) string {
	startingScriptTag := "<script"
	endingScriptTag := "</script>"

	var script string

	for {
		startingScriptTagIndex := strings.Index(s, startingScriptTag)
		endingScriptTagIndex := strings.Index(s, endingScriptTag)

		if startingScriptTagIndex > -1 && endingScriptTagIndex > -1 {
			script = s[startingScriptTagIndex : endingScriptTagIndex+len(endingScriptTag)]
			s = strings.Replace(s, script, "", 1)
			continue
		}

		break
	}

	return s
}

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

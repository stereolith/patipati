package main

import (
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getAttr(t html.Token, attr string) (href string) {
	// Iterate over token attributes until we find a specific attribute
	for _, a := range t.Attr {
		if a.Key == attr {
			href = a.Val
		}
	}
	return
}

// get website title
func getTitle(url string) string {
	resp, _ := http.Get(url)
	//bytes, _ := ioutil.ReadAll(resp.Body)

	z := html.NewTokenizer(resp.Body)

	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			return ""
		case tt == html.StartTagToken:
			t := z.Token()
			isTitle := t.Data == "title"
			if isTitle {
				text := z.Next()
				if text == html.TextToken {
					resp.Body.Close()
					return z.Token().Data
				}
			}
		}
	}

}

// get website title
func getIcon(link string) string {
	resp, _ := http.Get(link)
	//bytes, _ := ioutil.ReadAll(resp.Body)

	protocol := link[0 : strings.Index(link, "//")+2]
	requestUrl, _ := url.Parse(link)
	baseUrl := protocol + requestUrl.Hostname()

	var lastIcon string
	z := html.NewTokenizer(resp.Body)

	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			return lastIcon
		case tt == html.StartTagToken:
			t := z.Token()
			isLink := t.Data == "link"
			if isLink {
				switch rel := getAttr(t, "rel"); rel {
				case "apple-touch-icon":
					if href := getAttr(t, "href"); strings.Contains(href, "//") {
						return href
					} else {
						return baseUrl + href
					}
				case "icon":
					if href := getAttr(t, "href"); strings.Contains(href, "//") {
						lastIcon = href
					} else {
						lastIcon = baseUrl + href
					}
				}
			}
		}
	}

}

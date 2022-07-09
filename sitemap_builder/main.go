package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"link_parser"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"

func main() {
	urlForParse := flag.String("url", "https://gophercises.com", "the url that you want to build a sitemap for")
	flag.Parse()

	hrefs := bfc(*urlForParse, 3)
	urlsXml := urlXml{
		Xmlns: xmlns,
	}
	for _, v := range hrefs {
		urlsXml.Url = append(urlsXml.Url, loc{v})
	}

	output, err := xml.MarshalIndent(urlsXml, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Print(xml.Header)
	os.Stdout.Write(output)
}

type loc struct {
	Value string `xml:"loc"`
}
type urlXml struct {
	XMLName xml.Name `xml:"urlset"`
	Url     []loc    `xml:"url"`
	Xmlns   string   `xml:"xmlns,attr"`
}

func bfc(strUrl string, maxDepth int) []string {
	seen := make(map[string]struct{})
	var q map[string]struct{}
	nq := map[string]struct{}{
		strUrl: struct{}{},
	}

	for i := 0; i <= maxDepth; i++ {
		q, nq = nq, make(map[string]struct{})
		for url, _ := range q {
			if _, ok := seen[url]; ok {
				continue
			}
			seen[url] = struct{}{}
			for _, link := range getRefs(url) {
				nq[link] = struct{}{}
			}
		}
	}
	ret := make([]string, 0, len(seen))
	for url, _ := range seen {
		ret = append(ret, url)
	}
	return ret
}
func getRefs(stringUrl string) []string {
	resp, err := http.Get(stringUrl)
	if err != nil {
		panic("kek")
	}
	defer resp.Body.Close()

	reqUrl := resp.Request.URL
	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}

	base := baseUrl.String()

	return hrefs(resp.Body, base)
}

func hrefs(r io.Reader, base string) []string {
	parsedLinks, _ := link_parser.Parse(r)
	var links []string
	for _, l := range parsedLinks {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			links = append(links, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			links = append(links, l.Href)
		}
	}

	return filter(links, withPrefix(base))
}

func filter(hrefs []string, keepFN func(string) bool) []string {
	var links []string
	for _, v := range hrefs {
		if keepFN(v) {
			links = append(links, v)
		}
	}

	return links
}

func withPrefix(base string) func(string) bool {
	return func(link string) bool {
		return strings.HasPrefix(link, base)
	}
}

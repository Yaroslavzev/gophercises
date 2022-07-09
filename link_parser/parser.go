package link_parser

import (
	"golang.org/x/net/html"
	"io"
	"strings"
)

type Link struct {
	Href string
	Text string
}

func Parse(r io.Reader) ([]Link, error) {
	data, err := html.Parse(r)
	if err != nil {
		panic(err)
	}

	var nodes []*html.Node
	nodes = linkNodes(data, nodes)
	links := getLinks(nodes)

	return links, err
}

func getLinks(nodes []*html.Node) []Link {
	var links []Link

	for _, n := range nodes {
		var link Link
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				link.Href = attr.Val

			}
		}
		text := link_text(n)
		link.Text = text
		links = append(links, link)
	}

	return links
}

func link_text(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}
	var ret string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += link_text(c)
	}
	return strings.Join(strings.Fields(ret), " ")
}

func linkNodes(n *html.Node, nodes []*html.Node) []*html.Node {

	if n.Type == html.ElementNode && (n.Data == "a") {
		nodes = append(nodes, n)

		return nodes
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		nodes = linkNodes(c, nodes)
	}
	return nodes
}

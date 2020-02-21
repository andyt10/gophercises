package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/net/html"
)

type link struct {
	Href string
	Text string
}

func main() {
	file, _ := openSource("tests/test4.html")

	loadSource(file)
}

// Loading and opening the parsing file
func openSource(fileLoc string) (*io.Reader, error) {
	file, err := os.Open(fileLoc)

	if err != nil {
		fmt.Println("Unable to open file:", fileLoc, ":", err)
	}

	var reader io.Reader
	reader = file
	return &reader, nil

}

func loadSource(reader *io.Reader) []link {

	node, err := html.Parse(*reader)

	if err != nil {
		fmt.Println("Unable to parse HTML source:", err)
		os.Exit(1)
	}

	var linksList = make([]link, 0)
	linksList = recursiveParse(node, linksList)
	return linksList
}

func sanatiseLinktext(linkText string) string {
	noNewLines := strings.Replace(linkText, "\n", "", -1)
	trimmed := strings.Trim(noNewLines, " ")
	return trimmed
}

func handleANode(n *html.Node) link {

	var newLink link
	for _, a := range n.Attr {
		if a.Key == "href" {
			newLink = link{Href: a.Val, Text: sanatiseLinktext(n.FirstChild.Data)}
			break
		}
	}
	return newLink
}

func recursiveParse(n *html.Node, links []link) []link {

	if n.Type == html.ElementNode && n.Data == "a" {
		newLink := handleANode(n)
		links = append(links, newLink)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = recursiveParse(c, links)
	}

	return links
}

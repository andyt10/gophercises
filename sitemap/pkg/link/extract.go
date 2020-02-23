package link

import (
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/net/html"
)

type ExtractedLink struct {
	Href string
	Text string
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

func ExtractLinksFromString(stringBody string) []ExtractedLink {
	//_ := strings.NewReader(stringBody)
	//ExtractLinks(&reader)
	return nil
}

func ExtractLinks(reader *io.Reader) []ExtractedLink {

	node, err := html.Parse(*reader)

	if err != nil {
		fmt.Println("Unable to parse HTML source:", err)
		os.Exit(1)
	}

	var linksList = make([]ExtractedLink, 0)
	linksList = recursiveParse(node, linksList)
	return linksList
}

func sanatiseLinktext(linkText string) string {
	noNewLines := strings.Replace(linkText, "\n", "", -1)
	return noNewLines
}

func handleANode(n *html.Node) ExtractedLink {

	var newLink ExtractedLink
	for _, a := range n.Attr {
		if a.Key == "href" {
			newLink = ExtractedLink{Href: a.Val, Text: strings.Trim(getAText(n, ""), " ")}
			break
		}
	}
	return newLink
}

func getAText(n *html.Node, linkText string) string {
	if n.Type == html.TextNode {
		return sanatiseLinktext(n.Data)
	}

	if n.Type != html.ElementNode {
		return ""
	}

	var new string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		new += getAText(c, linkText)
	}

	return new
}

func recursiveParse(n *html.Node, links []ExtractedLink) []ExtractedLink {

	if n.Type == html.ElementNode && n.Data == "a" {
		newLink := handleANode(n)
		links = append(links, newLink)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = recursiveParse(c, links)
	}

	return links
}

package main

import (
	"cor_gophercises/sitemap/pkg/link"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type urlParts struct {
	proto    string
	domain   string
	resource string
}

const namespaceConst string = "http://www.sitemaps.org/schemas/sitemap/0.9"

type xmlUrl struct {
	XMLName xml.Name `xml:"url"`
	Loc     string   `xml:"loc"`
}

type urlSet struct {
	UrlSet []xmlUrl `xml:"url"`
	Xmlns  string   `xml:"xmlns,attr"`
}

//*********
// Main Funcs
//*********
func main() {
	doRun()
}

func doRun() {
	//link.OpenSource()
	//

	pageBody, getParseErr := getPage("http://dcbfthwkrvlmznxs.neverssl.com/online")

	if getParseErr != nil {
		fmt.Println("Error getting first page:", getParseErr)
		os.Exit(1)
	}

	//	"io/ioutil"
	//body, _ := ioutil.ReadAll(pageBody)

	links := link.ExtractLinks(&pageBody)

	fmt.Println(links)

	//site, depth := parseArgs()

	//fmt.Println("Mapping Site:", site, "To a depth of:", depth, "links")
}

func parseArgs() (string, int) {

	maxDepth := flag.Int("max-depth", 3, "Max number of links to follow in a site before stopping.")
	siteName := flag.String("site-name", "", "A site to create a map for.")
	flag.Parse()

	return *siteName, *maxDepth

}

// *********
// XML Stuff
// *********
func buildMapXml(links []link.ExtractedLink) ([]byte, error) {

	xmlUrls := make([]xmlUrl, len(links))

	for i, v := range links {
		xmlUrls[i].Loc = v.Href
	}

	xmlData := urlSet{UrlSet: xmlUrls, Xmlns: namespaceConst}

	output, err := xml.MarshalIndent(xmlData, "  ", "    ")
	if err != nil {
		fmt.Printf("Error Mashalling XML Site Map: %v\n", err)
		return nil, err
	}

	return output, nil

}

// ***********************
// URL/Address Formatting
// ***********************
func noramliseAddress(url string) urlParts {
	//lower case it all
	//take http/https if present, put in 'proto' value (take everything before ://)
	//take everything after ":// but before first /", put in domain value
	// take everything after/including first "/ and put in resource field"
	url = strings.ToLower(url)

	proto := ""
	maybeProto := strings.Split(url, "://")
	if len(maybeProto) > 1 {
		//proto is present
		proto = maybeProto[0]

	}

	domain := ""
	maybeDomain := strings.SplitN(maybeProto[len(maybeProto)-1], "/", 2)
	if maybeDomain[0] != "" || len(maybeDomain) > 1 {
		//remove www.
		domain = strings.Trim(maybeDomain[0], "www.")
	}

	resource := "/"
	if maybeDomain[0] == "" || len(maybeDomain) > 1 {
		resource = resource + maybeDomain[1]
	}

	return urlParts{proto: proto, domain: domain, resource: resource}

}

func isLinkSameWebsite(linkData urlParts, site urlParts) bool {
	//It's the same site IF
	// domain value in linkData is "" (relative link)
	// domain in linkData == domain in site

	if linkData.domain == site.domain {
		return true
	}

	if linkData.domain == "" {
		return true
	}

	return false
}

// **************
// Website Traversal
// **************
func getPage(pageUrl string) (io.Reader, error) {

	resp, err := http.Get(pageUrl)

	if err != nil {
		fmt.Println("Error GET-ing page:", pageUrl, ":", err)
		return nil, err
	}

	//defer resp.Body.Close()
	return resp.Body, nil
}

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
	rootUrl := "http://dcbfthwkrvlmznxs.neverssl.com/online"
	doRun(noramliseAddress(rootUrl))
}

func doRun(rootSite urlParts) []urlParts {

	pageBody, getParseErr := getPage(makeUrlString(rootSite))

	if getParseErr != nil {
		fmt.Println("Error getting first page:", getParseErr)
		os.Exit(1)
	}

	var siteMap []urlParts

	pageLinks := getLinksInPage(pageBody)

	//if val, ok := dict["foo"]; ok {
	for _, v := range pageLinks {
		if isLinkSameWebsite(v, rootSite) {
			fmt.Println("same", v)
		} else {
			fmt.Println("not same", v)
		}
	}

	return siteMap

}

func doRunAux() {

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
func buildSiteMapXml(links []string, shouldIndent bool) ([]byte, error) {

	xmlUrls := make([]xmlUrl, len(links))

	for i, v := range links {
		xmlUrls[i].Loc = v
	}

	xmlData := urlSet{UrlSet: xmlUrls, Xmlns: namespaceConst}
	fmt.Println(xmlData)

	var output []byte
	var marshalError error

	if shouldIndent {
		output, marshalError = xml.MarshalIndent(xmlData, "  ", "    ")
	} else {
		output, marshalError = xml.Marshal(xmlData)
	}

	if marshalError != nil {
		fmt.Printf("Error Mashalling XML Site Map: %v\n", marshalError)
		return nil, marshalError
	}

	return output, nil

}

// ***********************
// URL/Address Formatting
// ***********************
func makeUrlString(parts urlParts) string {
	formattedUrl := fmt.Sprintf("%v://%v%v", parts.proto, parts.domain, parts.resource)
	return formattedUrl
}

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

/* Determines if a link is for the same (current trawling) website.
* IF it is, return true
 */
func isLinkSameWebsite(linkData urlParts, site urlParts) bool {

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

func getLinksInPage(pageData io.Reader) []urlParts {

	links := link.ExtractLinks(&pageData)

	formattedLinks := make([]urlParts, len(links))

	for i, v := range links {
		formattedLinks[i] = noramliseAddress(v.Href)
	}

	return formattedLinks

}

func getPage(pageUrl string) (io.Reader, error) {

	resp, err := http.Get(pageUrl)

	if err != nil {
		fmt.Println("Error GET-ing page:", pageUrl, ":", err)
		return nil, err
	}

	//defer resp.Body.Close()
	return resp.Body, nil
}

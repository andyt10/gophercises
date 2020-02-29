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
	maxDepth := 2
	siteMap, err := doRun(noramliseAddress(rootUrl), maxDepth)

	if err != nil {
		fmt.Println("ERROR in SiteMap building:", err)
		os.Exit(1)
	}

	fmt.Println(string(siteMap))

}

func doRun(rootSite urlParts, maxDepth int) (string, error) {

	var siteMapData []string

	siteMapData = doRunAux(rootSite, rootSite, 1, maxDepth, siteMapData)
	siteMap, mapBuildErr := buildSiteMapXml(siteMapData, true)

	if mapBuildErr != nil {
		fmt.Println("Unable to build map for site data")
		return "", mapBuildErr
	}

	return string(siteMap), nil

}

func doRunAux(pageToGet urlParts, rootSite urlParts, currentDepth int, maxDepth int, siteMapData []string) []string {

	newDepth := currentDepth + 1

	if newDepth > maxDepth {
		fmt.Println("Already navigated as far as desired, returning existing sitemap")
		return siteMapData
	}

	if !isLinkSameWebsite(pageToGet, rootSite) {
		fmt.Println("Link is not for same site, not performing GET call")
		return siteMapData
	}

	//If already in sitemap, return.
	for _, v := range siteMapData {
		if makeUrlString(pageToGet) == v {
			fmt.Println("Already in map, returning")
			return siteMapData
		}
	}

	//Query site, and parse links. Returning sitemap
	pageBody, getParseErr := getPage(makeUrlString(rootSite))
	pageLinks := getLinksInPage(pageBody)

	if getParseErr != nil {
		fmt.Println("Error getting page:", makeUrlString(pageToGet), ":", getParseErr)
		return siteMapData
	}

	for _, v := range pageLinks {
		if isLinkSameWebsite(v, rootSite) {
			siteMapData = append(siteMapData, makeUrlString(v))
			siteMapData = doRunAux(v, rootSite, 1, maxDepth, siteMapData)

		}
	}

	return siteMapData

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

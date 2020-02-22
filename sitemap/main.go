package main

import (
	"cor_gophercises/sitemap/pkg/link"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {

}

type urlParts struct {
	proto    string
	domain   string
	resource string
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
	return true
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

func parseArgs() (string, int) {

	maxDepth := flag.Int("max-depth", 3, "Max number of links to follow in a site before stopping.")
	siteName := flag.String("site-name", "", "A site to create a map for.")
	flag.Parse()

	return *siteName, *maxDepth

}

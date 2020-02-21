package main

import (
	"flag"
	"fmt"
)

func main() {

	//	"cor_gophercises/sitemap/pkg/link"

	//link.OpenSource()

	site, depth := parseArgs()

	fmt.Println("Mapping Site:", site, "To a depth of:", depth, "links")
}

func parseArgs() (string, int) {

	maxDepth := flag.Int("max-depth", 3, "Max number of links to follow in a site before stopping.")
	siteName := flag.String("site-name", "", "A site to create a map for.")
	flag.Parse()

	return *siteName, *maxDepth

}

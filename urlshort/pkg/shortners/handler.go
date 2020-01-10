package shortners

import (
	"fmt"
	yaml "gopkg.in/yaml.v2"
	"net/http"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	//	TODO: Implement this...

	hello := func(w http.ResponseWriter, r *http.Request) {
		if dest, ok := pathsToUrls[r.URL.Path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
	return hello
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO: Implement this...
	err, parsedData := parseYaml(yml)

	if err != nil {
		fmt.Println("Error Parsing YAML")
		fmt.Println(err)
		return nil, err
	}
	mappy := buildMap(parsedData)

	return MapHandler(mappy, fallback), nil
}

func buildMap(data []yamlData) map[string]string {
	pathsToUrls := make(map[string]string)
	for _, pu := range data {
		pathsToUrls[pu.Path] = pu.Url
	}
	return pathsToUrls
}

func parseYaml(data []byte) (error, []yamlData) {
	var pathMap []yamlData
	err := yaml.Unmarshal(data, &pathMap)

	fmt.Println(pathMap)

	if err != nil {
		return err, nil
	}
	return nil, pathMap
}

type yamlData struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

//func hello(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintln(w, "Hello, world!")
//}

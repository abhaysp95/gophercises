package urlshort

import (
	"net/http"

	"gopkg.in/yaml.v2"
)

// MapHandler will return http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any path (keys in the map) to their corresponding
// URL (values that each key points to in the map in string format). If the
// path is not found, then fallback http.Handler will be called instead
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		path := req.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, req, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, req)
	}
}

// YamlHandler will parse the provided YAML and will return http.HandlerFunc
// (which also implements http.Handler) that will attempt to map any path to
// their corresponding URL. If the path is not found, then fallback
// http.Handler will be called instead
//
// YAML is expected to be in format:
//    - path: /some-path
//	    url: http://path-url
//
// The error returned will be related to YAML errors
func YamlHandler(yamlData []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathUrls, err := parseYaml(yamlData) // parse Yaml
	if err != nil {
		panic(err)
	}
	pathsToUrls := buildMap(pathUrls) // convert to map
	return MapHandler(pathsToUrls, fallback), nil
}

// yamlParser
func parseYaml(data []byte) ([]yamlStruct, error) {
	var pathUrls []yamlStruct
	if err := yaml.Unmarshal(data, &pathUrls); err != nil {
		return nil, err
	}
	return pathUrls, nil
}

// buildMap makes map from []yamlstruct
func buildMap(data []yamlStruct) map[string]string {
	m := make(map[string]string)
	for _, pu := range data {
		m[pu.Path] = pu.Url
	}
	return m
}

type yamlStruct struct {
	Path string `yaml:"path"`
	Url string `yaml:"url"`
}

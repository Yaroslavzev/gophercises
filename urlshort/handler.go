package urlshort

import (
	"encoding/json"
	yaml "gopkg.in/yaml.v3"
	"net/http"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
		}
		fallback.ServeHTTP(w, r)
	}
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
	var pathURLs []CommonStruct
	err := yaml.Unmarshal(yml, &pathURLs)
	if err != nil {
		return nil, err
	}
	pathMap := buildMap(pathURLs)

	return MapHandler(pathMap, fallback), err
}

func JSONHandler(jsonBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var pathURLs []CommonStruct
	err := json.Unmarshal(jsonBytes, &pathURLs)
	if err != nil {
		return nil, err
	}
	pathMap := buildMap(pathURLs)
	//fmt.Println(pathMap)

	return MapHandler(pathMap, fallback), err
}

type CommonStruct struct {
	Path string
	Url  string
}

//type JsonStruct struct {
//	Path string `json:"path"`
//	Url  string `json:"url"`
//}
//
//type YamlStruct struct {
//	Path string `yaml:"path"`
//	Url  string `yaml:"url"`
//}

func buildMap(pathURLs []CommonStruct) map[string]string {
	pathURLsMap := make(map[string]string)

	for _, v := range pathURLs {
		pathURLsMap[v.Path] = v.Url
	}

	return pathURLsMap
}

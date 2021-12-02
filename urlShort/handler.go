package urlshort

import (
	"net/http"

	yaml "gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {

	return func(rw http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if done, ok := pathsToUrls[path]; ok{
			http.Redirect(rw, r, done, http.StatusFound)
			return
		}
		fallback.ServeHTTP(rw,r)
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
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathUrls, err := parseYaml(yml)
	if err != nil{
		return nil, err
	}
	pathsToUrls := buildMap(pathUrls)
	return MapHandler(pathsToUrls, fallback), nil
}

func parseYaml(data []byte) ([]pathsUrl, error){
	var pathUrls []pathsUrl
	err := yaml.Unmarshal(data, &pathUrls)
	if err != nil{
		return nil, err
	}
	return pathUrls, nil
}

func buildMap(pathUrls []pathsUrl) map[string]string{
	pathsToUrls := make(map[string]string)
	for _, pu := range(pathUrls){
		pathsToUrls[pu.Path] = pu.Url
	}
	return pathsToUrls
}


type pathsUrl struct{
	Path string `yaml:"path"`
	Url string `yaml:"url"`
}
package urlshort

import (
	"net/http"
	"gopkg.in/yaml.v2"
	"fmt"
	
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	
	return func(w http.ResponseWriter, r *http.Request) {
		path:= r.URL.RequestURI()
		redirectedPath :=  pathsToUrls[path]

		if(redirectedPath != ""){
			http.Redirect(w,r,redirectedPath,http.StatusSeeOther)
		} else {
			fallback.ServeHTTP(w,r)
		}
    }
}

//InstanceConfig is the struct for the config
type InstanceConfig struct {
	Path []string
	URL []string
}

func parseYAML(yml []byte) (InstanceConfig,error) {
	var config InstanceConfig
	err := yaml.Unmarshal(yml, &config)
    if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return config,nil
}

func buildMap(instance InstanceConfig) map[string]string {
		var urlMap map[string]string
		urlMap =  make(map[string]string)
		fmt.Println("instance is",instance)
		for i := range instance.Path {
			urlMap[instance.Path[i]] = instance.URL[i]
		} 
		return urlMap
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
		parsedYaml, err := parseYAML(yml)
		if err != nil {
			return nil, err
		}
		fmt.Println(parsedYaml)
		pathMap := buildMap(parsedYaml)
		return MapHandler(pathMap, fallback), nil
}




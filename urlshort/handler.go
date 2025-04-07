package urlshort

import (
	"net/http"

	"gopkg.in/yaml.v3"
)

type Yaml struct {
	Path string `yaml:"path"`
	Url string `yaml:"url"`
}

func buildMap(parsedYaml []Yaml) map[string]string {
	pathMap := make(map[string]string) 
	for _, yml := range parsedYaml {
		pathMap[yml.Path] = yml.Url
	}
	return pathMap
}

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if url, ok := pathsToUrls[r.URL.Path]; ok {
			http.Redirect(w, r, url, http.StatusMovedPermanently)	
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

func ParseYaml(Data []byte) ([]Yaml, error) {
	var yamlData []Yaml	
	err := yaml.Unmarshal(Data, &yamlData)
	return yamlData, err
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYaml, err := ParseYaml(yml)
	CheckErr(err)

	pathMap := buildMap(parsedYaml)	
	return MapHandler(pathMap, fallback), nil
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
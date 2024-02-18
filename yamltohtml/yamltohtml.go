package yamltohtml

import (
	"bytes"
	"os"
	"text/template"

	"gopkg.in/yaml.v2"
)

type PageData struct {
	Title string `yaml:"Title"`
	Body  string `yaml:"Body"`
}

func Yamltohtml(path string) (string, error) {
	tmpl, err := template.New("page").Parse(`<html><head><title>{{.Title}}</title></head><body>{{.Body}}</body></html>`)
	if err != nil {
		return "", err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	var pd PageData
	err = yaml.Unmarshal(data, &pd)
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer

	if err := tmpl.Execute(&tpl, pd); err != nil {
		return "", err
	}

	return tpl.String(), nil
}

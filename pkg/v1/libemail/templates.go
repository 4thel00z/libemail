package libemail

import (
	"bytes"
	"html/template"
	"io"
	"io/ioutil"
)

type Template struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

func RenderTemplate(name, in string, m map[string]string) (string, error) {
	parsed, err := template.New(name).Parse(in)
	if err != nil {
		return "", err
	}
	buffer := bytes.NewBufferString("")
	err = parsed.Execute(buffer, m)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}

func RenderTemplateFromFs(path string, m map[string]string) (string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return RenderTemplate(path, string(content), m)
}

func RenderTemplateFromReader(name string, r io.Reader, m map[string]string) (string, error) {
	content, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	return RenderTemplate(name, string(content), m)
}

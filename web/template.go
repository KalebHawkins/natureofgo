package web

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
)

//go:embed templates/*
var TemplateFS embed.FS

type Template struct {
	HTMLTemplate *template.Template
}

func ParseTemplate(fs embed.FS, pattern ...string) (Template, error) {
	htmlTpl, err := template.ParseFS(fs, pattern...)

	if err != nil {
		return Template{}, fmt.Errorf("error parsing template: %w", err)
	}

	return Template{
		HTMLTemplate: htmlTpl,
	}, nil
}

func (t Template) Execute(w http.ResponseWriter, data any) {
	t.HTMLTemplate.Execute(w, data)
}

func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return t
}

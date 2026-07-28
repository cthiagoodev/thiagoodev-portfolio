package templates

import (
	"html/template"
	"io"

	"github.com/cthiagoodev/thiagoodev-portfolio/templates"
)

type TemplateManager struct {
	template *template.Template
}

func NewTemplateManager() (*TemplateManager, error) {
	paths := []string{
		"*.html",
		"components/*.html",
		"static/css/*.css",
	}

	result, err := template.ParseFS(templates.TemplateFS, paths...)

	if err != nil {
		return nil, err
	}

	return &TemplateManager{result}, nil
}

func (tm *TemplateManager) Render(w io.Writer, name string, data any) error {
	return tm.template.ExecuteTemplate(w, name, data)
}

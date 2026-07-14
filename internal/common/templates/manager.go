package templates

import (
	"embed"
	"html/template"
	"io"
)

var templateFS embed.FS

type TemplateManager struct {
	template *template.Template
}

func NewTemplateManager() (*TemplateManager, error) {
	result, err := template.ParseFS(templateFS, "templates/*.html", "templates/pages/*.html")

	if err != nil {
		return nil, err
	}

	return &TemplateManager{result}, nil
}

func (tm *TemplateManager) Render(w io.Writer, name string, data any) error {
	return tm.template.ExecuteTemplate(w, name, data)
}

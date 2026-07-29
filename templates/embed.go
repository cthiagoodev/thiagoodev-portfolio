package templates

import "embed"

//go:embed *.html components/*.html
var TemplateFS embed.FS

//go:embed static
var StaticsFS embed.FS

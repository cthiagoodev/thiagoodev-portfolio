package templates

import "embed"

//go:embed *.html pages/*.html components/*.html
var TemplateFS embed.FS

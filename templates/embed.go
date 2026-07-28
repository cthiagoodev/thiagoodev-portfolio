package templates

import "embed"

//go:embed *.html components/*.html static/css/*.css
var TemplateFS embed.FS

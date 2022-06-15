package config

import (
	"fmt"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

//Template html renderer
type Template struct {
	templates *template.Template
}

//Render Template echo html renderer
func (t *Template) RenderRaw(w io.Writer, name string, data interface{}) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func (t *Template) Render(w io.Writer, name string, data interface{}, context echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Renderer() *Template {
	templatesPath := fmt.Sprintf("%s/*.html", appConfig.Server.TemplatesPath)
	CachedTemplates := template.Must(template.ParseGlob(templatesPath))

	t := &Template{
		templates: CachedTemplates,
	}

	return t
}

package web

import (
	"embed"
	"fmt"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

var (
	//go:embed *.htm
	htmFiles embed.FS

	//go:embed dist
	Dist embed.FS
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

var (
	Templates = &Template{
		templates: template.Must(template.ParseFS(htmFiles, "*.htm")),
	}
)

// returns a function that will check if *T is nil, if it is, it will return nilStr.
// otherwise it will fmt.Sprintf(fmt, *T)
func PointerToStr[T any](fmtStr string, nilStr string) func(*T) string {
	return func(t *T) string {
		if t == nil {
			return nilStr
		}
		return fmt.Sprintf(fmtStr, *t)
	}
}

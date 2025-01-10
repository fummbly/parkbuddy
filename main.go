package main

import (
	"encoding/xml"
	"io"
	"log"
	"text/template"

	xmlHelper "github.com/fummbly/parkbuddy/internal/xml"
	"github.com/labstack/echo/v4"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

func main() {

	e := echo.New()

	e.Renderer = newTemplate()

	e.Static("/css", "css")
	e.Static("/scripts", "scripts")

	data, err := xmlHelper.ReadFile()

	if err != nil {
		log.Fatalf("Failed to read file: %v\n", err)
		return
	}

	v := xmlHelper.OSM{}

	err = xml.Unmarshal(data, &v)
	if err != nil {
		log.Fatalf("Failed to unmarshal xml: %v\n", err)
		return
	}

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", nil)
	})
	e.Logger.Fatal(e.Start(":8080"))
}

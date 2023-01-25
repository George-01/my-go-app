package render

import (
	"fmt"
	"net/http"
	"html/template"
)

//RenderTemplate renders templates using html-template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplates, _ := template.ParseFiles("./templates/" + tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplates.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing templates:", err)
		return
	}
}
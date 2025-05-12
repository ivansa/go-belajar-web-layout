package handler

import (
	"net/http"
	"text/template"
)

func TemplateLayout(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"./html-layout-file/header.gohtml",
		"./html-layout-file/footer.gohtml",
		"./html-layout-file/layout.gohtml",
	))

	t.ExecuteTemplate(w, "layout.gohtml", map[string]interface{}{
		"Title":   "My Web Page",
		"Name":    "John Doe",
		"Content": "This is the content of the page.",
	})
}

func TemplateLayoutDefine(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"./html-layout-define/header.gohtml",
		"./html-layout-define/footer.gohtml",
		"./html-layout-define/layout.gohtml",
	))

	t.ExecuteTemplate(w, "layout", map[string]interface{}{
		"Title":   "My Web Page",
		"Name":    "John Doe",
		"Content": "This is the content of the page.",
	})
}

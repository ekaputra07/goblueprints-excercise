package main

import (
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

func render(filename string) http.HandlerFunc {
	var tpl *template.Template
	var once sync.Once
	once.Do(func() {
		tpl = template.Must(template.ParseFiles(filepath.Join("templates", filename)))
	})
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, r) // we pass Request as the data so we can access Request data in template
	}
}
package main

import (
	"html/template"
	"net/http"
	"strconv"
)

func fizzbuzzHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/fizzbuzz.html"))

	if r.Method == http.MethodPost {
		numStr := r.FormValue("number")
		if num, err := strconv.Atoi(numStr); err == nil {
			tmpl.Execute(w, num)
			return
		}
	}

	tmpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", fizzbuzzHandler)
	http.ListenAndServe(":8080", nil)
}

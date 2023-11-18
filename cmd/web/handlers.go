package main

import (
	"html/template"
	"net/http"
)

type Post struct {
	UserName string
	UserImg  string
	Title    string
	PostImg  string
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	files := []string{
		"./ui/html/pages/base.html",
		"./ui/html/partials/header.html",
		"./ui/html/partials/nav.html",
		"./ui/html/partials/post.html",
		"./ui/html/pages/home.html",
	}

	posts := map[string][]Post{
		"Posts": {
			{UserName: "catuser", UserImg: "cat.jpeg", Title: "cat loves this food", PostImg: "foodpic.jpeg"},
			{UserName: "kewlcat", UserImg: "cat.jpeg", Title: "im a kewl cat", PostImg: "recipe2.jpeg"},
		},
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}
	err = ts.ExecuteTemplate(w, "base", posts)
	if err != nil {
		app.serverError(w, err)
	}
}

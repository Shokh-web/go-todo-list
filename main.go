// -*- Author: real0x0a1 -*-
// -*- Info: A lightweight and easy-to-use Todo List app -*-
// -*- File: main. -*-


package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var todos = []Todo{}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/create", createTodo)
	http.HandleFunc("/delete/", deleteTodo)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, todos)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	if title == "" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	todo := Todo{ID: len(todos) + 1, Title: title, Done: false}
	todos = append(todos, todo)
	http.Redirect(w, r, "/", http.StatusFound)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/delete/"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
	}
	http.Error(w, "Todo not found", http.StatusNotFound)
}

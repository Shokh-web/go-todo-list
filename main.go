// -*- Author: real0x0a1 -*-
// -*- Info: A lightweight and easy-to-use Todo List app -*-
// -*- File: main.go -*-


package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/Shokh-web/todo-app/database"
)

func main() {
	// Initialize database connection
	database.InitDB()

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/create", createTodo)
	http.HandleFunc("/delete/", deleteTodo)
	http.HandleFunc("/toggle/", toggleTodo)

	log.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var todos []database.Todo
	result := database.DB.Find(&todos)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, todos)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	title := r.FormValue("title")
	description := r.FormValue("description")
	if title == "" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	todo := database.Todo{
		Title:       title,
		Description: description,
	}

	result := database.DB.Create(&todo)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/delete/")
	result := database.DB.Delete(&database.Todo{}, id)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func toggleTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/toggle/"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var todo database.Todo
	result := database.DB.First(&todo, id)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}

	todo.Completed = !todo.Completed
	result = database.DB.Save(&todo)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Todo list of activity
type Todo struct {
	Title string
	Done  bool
}

// TodoPageData list of page
type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	tmpl := template.Must(template.ParseFiles("layout.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My TODO list",
			Todos: []Todo{
				{Title: "Task 1 is here", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3 last", Done: true},
			},
		}
		tmpl.Execute(w, data)
	})

	fmt.Println("Listening on localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

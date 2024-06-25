package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ToDo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ToDoList []ToDo

var toDoList ToDoList

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		getToDoListHandler(w, r)

	})

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}

func getToDoListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(toDoList)
}

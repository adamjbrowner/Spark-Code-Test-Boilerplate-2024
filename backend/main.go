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

var toDoList = ToDoList{}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// naughty CORS, only for development
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		switch r.Method {
		case "GET":
			getToDoListHandler(w)
		case "POST":
			createToDoHandler(w, r)
		case "OPTIONS":
			w.WriteHeader(http.StatusOK)
		}
	})

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}

func getToDoListHandler(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(toDoList)
}

func createToDoHandler(w http.ResponseWriter, r *http.Request) {
	var newToDo ToDo

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&newToDo)

	if err != nil {
		http.Error(w, "Invalid Input", http.StatusBadRequest)
		return
	}

	toDoList = append(toDoList, newToDo)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newToDo)
}

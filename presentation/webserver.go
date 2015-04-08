package main

import (
	"encoding/json"
	"net/http"
)

type Todo struct {
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

type Todos []Todo

func index(w http.ResponseWriter, r *http.Request) {
	todos := []Todo{Todo{Name: "write presentation"}, Todo{Name: "host meetup"}}
	json.NewEncoder(w).Encode(todos)
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":9000", nil)
}

package main

import (
	"net/http"

	"github.com/SuperTikuwa/golang-handson-todo/handler"
)

func main() {
	http.HandleFunc("/task", handler.TaskHandler)

	http.ListenAndServe(":8081", nil)
}

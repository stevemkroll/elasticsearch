package api

import (
	"elasticsearch/employee"
	"elasticsearch/task"
	"net/http"
	"os"
)

func Run() {
	http.HandleFunc("/employee", employee.SearchHandler)
	http.HandleFunc("/task", task.SearchHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		os.Exit(1)
	}
}

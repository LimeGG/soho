package web

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Task struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var tasks = make(map[string]Task)

// Проверка выполнения задания
func checkTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID := vars["taskID"]
	task, exists := tasks[taskID]
	if !exists {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	// Логика для проверки выполнения задания
	task.Status = "completed"
	tasks[taskID] = task
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func Task_main() {
	r := mux.NewRouter()
	r.HandleFunc("/task/{taskID}/complete", checkTask).Methods("POST")
	log.Fatal(http.ListenAndServe(":8002", r))
}

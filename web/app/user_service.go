package web

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type User struct {
	ID      string `json:"id"`
	Balance int    `json:"balance"`
}

var users = make(map[string]User)

// Обновление баланса пользователя
func updateBalance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userID"]
	// Для простоты увеличим баланс на 10 монет
	user, exists := users[userID]
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	user.Balance += 10
	users[userID] = user
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func User_main() {
	r := mux.NewRouter()
	r.HandleFunc("/user/{userID}/balance", updateBalance).Methods("POST")
	log.Fatal(http.ListenAndServe(":8001", r))
}

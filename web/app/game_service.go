package web

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Game struct {
	ID     string  `json:"id"`
	Board  [][]int `json:"board"`
	UserID string  `json:"user_id"`
	Status string  `json:"status"` // e.g., "in_progress", "finished"
}

var games = make(map[string]Game)

// Генерация игрового поля (5x5) с случайными числами для примера
func generateBoard() [][]int {
	board := make([][]int, 5)
	for i := range board {
		board[i] = make([]int, 5)
		for j := range board[i] {
			board[i][j] = (i + j) % 3 // Просто пример для генерации значений
		}
	}
	return board
}

// Старт новой игры
func startGame(w http.ResponseWriter, r *http.Request) {
	game := Game{
		ID:     "12345",
		Board:  generateBoard(),
		UserID: "user1", // Должен быть ID из JWT или внешней системы
		Status: "in_progress",
	}
	games[game.ID] = game
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(game)
}

// Открытие клетки
func openCell(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameID := vars["gameID"]
	// Проверяем, существует ли игра
	game, exists := games[gameID]
	if !exists {
		http.Error(w, "Game not found", http.StatusNotFound)
		return
	}
	// Логика открытия клетки и начисления бонусов
	game.Status = "finished" // Пример, что игра завершена
	games[gameID] = game
	fmt.Fprintf(w, "Game %s finished.", gameID)
}

func Game_main() {
	r := mux.NewRouter()
	r.HandleFunc("/start", startGame).Methods("POST")
	r.HandleFunc("/game/{gameID}/open", openCell).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", r))
}

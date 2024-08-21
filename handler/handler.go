package handler

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"sync"
	"test-task/entity"
	"test-task/storage"
)

var (
	mu    sync.Mutex
	store = storage.NewFileStorage("operations.json")
)

func OperationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var op entity.Operation
	if err := json.NewDecoder(r.Body).Decode(&op); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if op.Cmd == "" || op.ID == 0 {
		http.Error(w, "Invalid cmd or id", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	if result, found := store.Get(op.ID); found {
		jsonResponse(w, op, result)
		return
	}

	result := generateRandomResult()
	store.Save(op.ID, result)
	jsonResponse(w, op, result)
}

func jsonResponse(w http.ResponseWriter, op entity.Operation, result string) {
	response := map[string]interface{}{
		"cmd":    op.Cmd,
		"id":     op.ID,
		"result": result,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func generateRandomResult() string {
	results := []string{"red", "blue", "green", "yellow", "purple", "orange", "black", "white"}
	return results[rand.Intn(len(results))]
}

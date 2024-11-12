package server

import (
	"encoding/json"
	"fmt"
	"github.com/erik-olsson-op/shared/database"
	"github.com/erik-olsson-op/shared/logger"
	"net/http"
	"sync"
)

func Init(port string, wg *sync.WaitGroup) {
	defer wg.Done()
	http.HandleFunc("/consume", consumerHandler)
	http.HandleFunc("/health", healthHandler)
	addr := fmt.Sprintf(":%v", port)
	logger.Logger.Infof("HTTP server is running on port %v", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		logger.Logger.Fatal(err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method only GET", http.StatusMethodNotAllowed)
		return
	}
	logger.Logger.Info("/health was requested!")
	w.WriteHeader(http.StatusOK)
}

func consumerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method only GET", http.StatusMethodNotAllowed)
		return
	}
	logger.Logger.Info("/consume was requested!")
	persons := database.Read()
	jsonData, err := json.Marshal(&persons)
	if err != nil {
		logger.Logger.Error(err)
		http.Error(w, "Failed", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonData)
	if err != nil {
		logger.Logger.Error(err)
		http.Error(w, fmt.Sprintf("ERR: %v", err), http.StatusBadRequest)
		return
	}
}

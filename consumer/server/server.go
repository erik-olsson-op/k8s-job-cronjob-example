package server

import (
	"encoding/json"
	"fmt"
	"github.com/erik-olsson-op/shared/database"
	"github.com/erik-olsson-op/shared/logger"
	"github.com/erik-olsson-op/shared/utils"
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
	persons := database.Read()
	logger.Logger.Info(fmt.Sprintf("/consume was requested, found %d persons in db", len(persons)))
	jsonData, err := json.Marshal(&persons)
	if err != nil {
		logger.Logger.Error(err)
		http.Error(w, "Failed", http.StatusInternalServerError)
		return
	}
	w.Header().Add(utils.HeaderContentType, "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonData)
	if err != nil {
		logger.Logger.Error(err)
		http.Error(w, fmt.Sprintf("ERR: %v", err), http.StatusBadRequest)
		return
	}
}

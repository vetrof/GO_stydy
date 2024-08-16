package handler

import (
	"apikafka/models" // Импортируем пакет моделей
	"encoding/json"
	"net/http"
)

func ProductsHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(response).Encode(models.GetProducts())
	if err != nil {
		http.Error(response, "Не удалось закодировать продукты", http.StatusInternalServerError)
	}
}

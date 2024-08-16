package router

import (
	"apikafka/handler"
	"github.com/gorilla/mux" // Используем популярный пакет для маршрутизации (необязательно)
)

// SetupRouter инициализирует маршрутизатор и связывает маршруты с хендлерами
func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	// Определяем маршруты и связываем их с хендлерами
	router.HandleFunc("/api/v1/products", handler.ProductsHandler).Methods("GET")

	return router
}

package main

import (
	"apikafka/router"
	"net/http"
)

func main() {

	// Пример отправки сообщения в Kafka
	//go kafka.ProduceMessage("my-topic", "Привет, Kafka!")

	// Пример получения сообщений из Kafka
	//go kafka.ConsumeMessages("my-topic")

	// Создаем маршрутизатор и связываем маршруты с хендлерами
	router := router.SetupRouter()

	// Запускаем сервер
	http.ListenAndServe(":8080", router)
}

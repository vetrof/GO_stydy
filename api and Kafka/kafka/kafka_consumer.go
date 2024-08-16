package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
)

func ConsumeMessages(topic string) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"}, // Замените на ваш адрес брокера
		Topic:   topic,
		GroupID: "my-group", // Уникальный идентификатор группы консьюмеров
	})
	defer reader.Close()

	for {
		message, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal("Ошибка чтения сообщения:", err)
		}
		log.Printf("Получено сообщение: %s\n", string(message.Value))
	}
}

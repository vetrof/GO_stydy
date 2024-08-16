package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
)

func ProduceMessage(topic, message string) {
	// Укажите адрес Kafka-брокера
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"}, // Замените на ваш адрес брокера
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
	defer writer.Close()

	err := writer.WriteMessages(context.Background(),
		kafka.Message{
			Value: []byte(message),
		},
	)

	if err != nil {
		log.Fatal("Не удалось отправить сообщение:", err)
	}
	log.Println("Сообщение отправлено:", message)
}

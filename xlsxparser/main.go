package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"log"
	"net/http"
	"os"

	"google.golang.org/api/sheets/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	spreadsheetID := os.Getenv("spreadsheetID")
	apiKey := os.Getenv("GOOGLE_SHEETS_API_KEY")

	// Инициализация HTTP клиента с использованием API ключа
	client := &http.Client{
		Transport: &transportAPIKey{Key: apiKey},
	}

	// Инициализация клиента Google Sheets API
	ctx := context.Background()
	service, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	// Запрос данных из таблицы
	readRange := "Лист1!A:B"
	resp, err := service.Spreadsheets.Values.Get(spreadsheetID, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	// Обработка полученных данных
	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
	} else {
		fmt.Println("Name, Major:")
		for _, row := range resp.Values {
			// Печать значений из таблицы
			fmt.Printf("%s, %s\n", row[0], row[1])
		}
	}
}

// transportAPIKey позволяет использовать API ключ в запросах
type transportAPIKey struct {
	Key string
}

func (t *transportAPIKey) RoundTrip(req *http.Request) (*http.Response, error) {
	// Добавляем API ключ в запрос
	req.URL.RawQuery += "&key=" + t.Key
	return http.DefaultTransport.RoundTrip(req)
}

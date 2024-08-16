package models

// Product представляет структуру данных продукта
type Product struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// GetProducts возвращает список продуктов (может быть реализовано с подключением к БД)
func GetProducts() []Product {
	return []Product{
		{Name: "vasya", Description: "looser"},
		{Name: "vasya2", Description: "looser2"},
	}
}

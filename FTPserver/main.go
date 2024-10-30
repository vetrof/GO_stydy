package main

import (
	"goftp.io/server/core"
	"goftp.io/server/driver/file"
	"log"
)

// Структура аутентификации
type MyAuth struct{}

func (a *MyAuth) CheckPasswd(username, password string) (bool, error) {
	// Проверка имени пользователя и пароля
	if username == "user" && password == "password" {
		return true, nil
	}
	return false, nil
}

func main() {
	// Создаем экземпляр драйвера файловой системы
	factory := &file.DriverFactory{
		RootPath: "/Users/vetrof/Documents/CODE/GO/go_study/FTPserver/ftp-root/", // Укажите путь к директории
	}

	// Конфигурация сервера
	config := &core.ServerOpts{
		Factory:      factory,
		Auth:         &MyAuth{},
		Port:         21,            // Порт для сервера
		PassivePorts: "30000-30100", // Диапазон пассивных портов
	}

	// Запуск FTP-сервера
	ftpServer := core.NewServer(config)
	log.Println("Starting FTP server...")

	if err := ftpServer.ListenAndServe(); err != nil {
		log.Fatal("Error starting FTP server:", err)
	}
}

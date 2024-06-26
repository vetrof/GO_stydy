package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// LogTime измеряет и выводит время выполнения функции
func LogTime(name string, start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}

func main() {
	start := time.Now() // Начало измерения времени выполнения main
	defer LogTime("main", start)

	// Открываем исходный файл для чтения
	inputFilePath := "C:\\CODE\\GO\\GO_stydy\\rebrain\\module02-defer\\06_task_in.txt"
	file, err := os.Open(inputFilePath)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer file.Close()

	// Создаем (или открываем) целевой файл для записи
	outputFilePath := "C:\\CODE\\GO\\GO_stydy\\rebrain\\module02-defer\\out.txt"
	outFile, err := os.Create(outputFilePath)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outFile.Close()

	// Создаем сканер для чтения из исходного файла
	scanner := bufio.NewScanner(file)
	// Создаем писатель для записи в целевой файл
	writer := bufio.NewWriter(outFile)

	// Считываем строки из исходного файла и записываем их в целевой файл
	for scanner.Scan() {
		line := scanner.Text()
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Error writing to output file:", err)
			return
		}
	}

	// Проверяем на наличие ошибок при сканировании
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from input file:", err)
	}

	// Обязательно сбрасываем буфер для записи всех данных в файл
	writer.Flush()
}

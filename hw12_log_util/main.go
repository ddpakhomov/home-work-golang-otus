package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Определение флагов командной строки
	fileFlag := flag.String("file", "", "Путь к файлу логов")
	levelFlag := flag.String("level", "INFO", "Уровень логов для анализа")
	outputFlag := flag.String("output", "", "Путь к выходному файлу для статистики")

	// Парсинг флагов командной строки
	flag.Parse()

	// Получение значений флагов или переменных окружения
	logFile := getEnvOrFlag("LOG_ANALYZER_FILE", *fileFlag)
	logLevel := getEnvOrFlag("LOG_ANALYZER_LEVEL", *levelFlag)
	outputFile := getEnvOrFlag("LOG_ANALYZER_OUTPUT", *outputFlag)

	// Проверка обязательного флага -file
	if logFile == "" {
		fmt.Println(
			"Ошибка: путь к файлу логов должен быть указан с помощью флага -file или переменной окружения LOG_ANALYZER_FILE",
		)
		os.Exit(1)
	}

	// Открытие лог-файла
	file, err := os.Open(logFile)
	if err != nil {
		fmt.Printf("Ошибка: не удалось открыть файл логов: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// Анализ лог-файла
	stats, err := analyzeLogFile(file, logLevel)
	if err != nil {
		fmt.Printf("Ошибка: не удалось проанализировать файл логов: %v\n", err)
		return
	}

	// Вывод статистики
	if outputFile == "" {
		// Вывод в стандартный поток вывода
		fmt.Println(stats)
	} else {
		// Запись в указанный файл
		err = os.WriteFile(outputFile, []byte(stats), 0o600)
		if err != nil {
			fmt.Printf("Ошибка: не удалось записать в выходной файл: %v\n", err)
			return
		}
	}
}

func getEnvOrFlag(envVar, flagValue string) string {
	if flagValue != "" {
		return flagValue
	}
	return os.Getenv(envVar)
}

func analyzeLogFile(file *os.File, logLevel string) (string, error) {
	scanner := bufio.NewScanner(file)
	levelCount := 0
	totalCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		totalCount++
		if strings.Contains(line, logLevel) {
			levelCount++
		}
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("не удалось прочитать файл логов: %w", err)
	}

	stats := fmt.Sprintf(
		"Всего записей логов: %d\nЗаписей логов с уровнем '%s': %d\n",
		totalCount, logLevel, levelCount,
	)
	return stats, nil
}

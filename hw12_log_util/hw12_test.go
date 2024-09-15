package main

import (
	"os"
	"testing"
)

func TestGetEnvOrFlag(t *testing.T) {
	// Тестовый случай: флаг задан
	flagValue := "flagValue"
	envVar := "TEST_ENV_VAR"
	expected := flagValue
	result := getEnvOrFlag(envVar, flagValue)
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}

	// Тестовый случай: переменная окружения задана
	os.Setenv(envVar, "envValue")
	flagValue = ""
	expected = "envValue"
	result = getEnvOrFlag(envVar, flagValue)
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}

	// Тестовый случай: ни флаг, ни переменная окружения не заданы
	os.Unsetenv(envVar)
	expected = ""
	result = getEnvOrFlag(envVar, flagValue)
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestAnalyzeLogFile(t *testing.T) {
	// Создание временного файла с тестовыми данными
	logData := `INFO: This is an info message
ERROR: This is an error message
INFO: Another info message
DEBUG: Debugging message
`
	file, err := os.CreateTemp("", "logfile")
	if err != nil {
		t.Fatalf("не удалось создать временный файл: %v", err)
	}
	defer os.Remove(file.Name())

	_, err = file.WriteString(logData)
	if err != nil {
		t.Fatalf("не удалось записать во временный файл: %v", err)
	}
	file.Close()

	// Открытие файла для чтения
	file, err = os.Open(file.Name())
	if err != nil {
		t.Fatalf("не удалось открыть временный файл: %v", err)
	}
	defer file.Close()

	// Тестовый случай: уровень логов INFO
	logLevel := "INFO"
	expected := "Всего записей логов: 4\nЗаписей логов с уровнем 'INFO': 2\n"
	result, err := analyzeLogFile(file, logLevel)
	if err != nil {
		t.Fatalf("не удалось проанализировать файл логов: %v", err)
	}
	if result != expected {
		t.Errorf("ожидалось %s, получено %s", expected, result)
	}

	// Перемотка файла в начало для следующего теста
	file.Seek(0, 0)

	// Тестовый случай: уровень логов ERROR
	logLevel = "ERROR"
	expected = "Всего записей логов: 4\nЗаписей логов с уровнем 'ERROR': 1\n"
	result, err = analyzeLogFile(file, logLevel)
	if err != nil {
		t.Fatalf("не удалось проанализировать файл логов: %v", err)
	}
	if result != expected {
		t.Errorf("ожидалось %s, получено %s", expected, result)
	}
}

package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"bin/bins"
)

// Create сохраняет список bin'ов в JSON файл
func Create(path string, binList *bins.BinList) error {
	// Создаем директорию если не существует
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("ошибка создания директории: %w", err)
	}

	// Сериализуем в JSON
	data, err := json.MarshalIndent(binList, "", "  ")
	if err != nil {
		return fmt.Errorf("ошибка сериализации JSON: %w", err)
	}

	// Записываем в файл
	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("ошибка записи файла: %w", err)
	}

	return nil
}

// Read читает список bin'ов из JSON файла
func Read(path string) (*bins.BinList, error) {
	// Проверяем существование файла
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// Если файл не существует, возвращаем пустой список
		return bins.NewBinList(), nil
	}

	// Читаем файл
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения файла: %w", err)
	}

	// Десериализуем JSON
	var binList bins.BinList
	if err := json.Unmarshal(data, &binList); err != nil {
		return nil, fmt.Errorf("ошибка десериализации JSON: %w", err)
	}

	return &binList, nil
}
package file

import (
	"os"
	"path/filepath"
	"strings"
)

func Create(path string) {
}

// Read читает содержимое файла и возвращает его как строку
func Read(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

// IsJSON проверяет, является ли файл JSON файлом по расширению
func IsJSON(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	return ext == ".json"
}
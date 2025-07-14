package file

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

func Create(path string) {
}

// Read читает содержимое файла и возвращает его как строку
func Read(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
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
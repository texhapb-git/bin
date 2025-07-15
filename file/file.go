package file

import (
	"os"
	"path/filepath"
	"strings"
)

// FileService интерфейс для работы с файлами
type FileService interface {
	Create(path string) error
	Read(path string) (string, error)
	Write(path string, content string) error
	IsJSON(path string) bool
	Exists(path string) bool
	Delete(path string) error
}

// FileManager реализация FileService
type FileManager struct{}

// NewFileManager создает новый экземпляр FileManager
func NewFileManager() FileService {
	return &FileManager{}
}

// Create создает файл по указанному пути
func (fm *FileManager) Create(path string) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	
	return nil
}

// Read читает содержимое файла и возвращает его как строку
func (fm *FileManager) Read(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

// Write записывает содержимое в файл
func (fm *FileManager) Write(path string, content string) error {
	return os.WriteFile(path, []byte(content), 0644)
}

// IsJSON проверяет, является ли файл JSON файлом по расширению
func (fm *FileManager) IsJSON(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	return ext == ".json"
}

// Exists проверяет существование файла
func (fm *FileManager) Exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// Delete удаляет файл
func (fm *FileManager) Delete(path string) error {
	return os.Remove(path)
}
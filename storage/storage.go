package storage

import (
	"encoding/json"
	"fmt"

	"bin/bins"
	"bin/file"
)

// StorageService интерфейс для работы с хранилищем
type StorageService interface {
	Save(binList *bins.BinList) error
	Load() (*bins.BinList, error)
	AddBin(bin *bins.Bin) error
	GetBin(id string) (*bins.Bin, error)
	GetAllBins() (*bins.BinList, error)
}

// JSONStorage реализация StorageService для JSON файлов
type JSONStorage struct {
	fileService file.FileService
	filePath    string
}

// NewJSONStorage создает новый экземпляр JSONStorage
func NewJSONStorage(fileService file.FileService, filePath string) StorageService {
	return &JSONStorage{
		fileService: fileService,
		filePath:    filePath,
	}
}

// Save сохраняет список bin'ов в JSON файл
func (js *JSONStorage) Save(binList *bins.BinList) error {
	// Создаем файл если не существует
	if err := js.fileService.Create(js.filePath); err != nil {
		return fmt.Errorf("ошибка создания файла: %w", err)
	}

	// Сериализуем в JSON
	data, err := json.MarshalIndent(binList, "", "  ")
	if err != nil {
		return fmt.Errorf("ошибка сериализации JSON: %w", err)
	}

	// Записываем в файл
	if err := js.fileService.Write(js.filePath, string(data)); err != nil {
		return fmt.Errorf("ошибка записи файла: %w", err)
	}

	return nil
}

// Load читает список bin'ов из JSON файла
func (js *JSONStorage) Load() (*bins.BinList, error) {
	// Проверяем существование файла
	if !js.fileService.Exists(js.filePath) {
		// Если файл не существует, возвращаем пустой список
		return bins.NewBinList(), nil
	}

	// Читаем файл
	data, err := js.fileService.Read(js.filePath)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения файла: %w", err)
	}

	// Десериализуем JSON
	var binList bins.BinList
	if err := json.Unmarshal([]byte(data), &binList); err != nil {
		return nil, fmt.Errorf("ошибка десериализации JSON: %w", err)
	}

	return &binList, nil
}

// AddBin добавляет bin в хранилище
func (js *JSONStorage) AddBin(bin *bins.Bin) error {
	binList, err := js.Load()
	if err != nil {
		return err
	}

	binList.AddBin(bin)
	return js.Save(binList)
}

// GetBin получает bin по ID
func (js *JSONStorage) GetBin(id string) (*bins.Bin, error) {
	binList, err := js.Load()
	if err != nil {
		return nil, err
	}

	for _, bin := range binList.Bins {
		if bin.ID == id {
			return &bin, nil
		}
	}

	return nil, fmt.Errorf("bin с ID %s не найден", id)
}

// GetAllBins получает все bins
func (js *JSONStorage) GetAllBins() (*bins.BinList, error) {
	return js.Load()
}
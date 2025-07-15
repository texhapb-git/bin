package api

import (
	"fmt"
	"time"

	"bin/bins"
	"bin/storage"
)

// BinService интерфейс для работы с bin'ами через API
type BinService interface {
	CreateBin(name string, private bool) (*bins.Bin, error)
	GetBin(id string) (*bins.Bin, error)
	GetAllBins() (*bins.BinList, error)
	DeleteBin(id string) error
}

// BinAPI реализация BinService
type BinAPI struct {
	storageService storage.StorageService
	key            string
}

// NewBinAPI создает новый экземпляр BinAPI
func NewBinAPI(storageService storage.StorageService, key string) BinService {
	return &BinAPI{
		storageService: storageService,
		key:            key,
	}
}

// CreateBin создает новый bin
func (ba *BinAPI) CreateBin(name string, private bool) (*bins.Bin, error) {
	// Генерируем уникальный ID (в реальном приложении здесь была бы более сложная логика)
	id := fmt.Sprintf("bin_%d", time.Now().Unix())
	
	bin := bins.NewBin(id, name, private)
	
	if err := ba.storageService.AddBin(bin); err != nil {
		return nil, fmt.Errorf("ошибка создания bin: %w", err)
	}
	
	return bin, nil
}

// GetBin получает bin по ID
func (ba *BinAPI) GetBin(id string) (*bins.Bin, error) {
	return ba.storageService.GetBin(id)
}

// GetAllBins получает все bins
func (ba *BinAPI) GetAllBins() (*bins.BinList, error) {
	return ba.storageService.GetAllBins()
}

// DeleteBin удаляет bin по ID
func (ba *BinAPI) DeleteBin(id string) error {
	// В простой реализации просто возвращаем ошибку
	// В реальном приложении здесь была бы логика удаления
	return fmt.Errorf("удаление bin'ов пока не реализовано")
}
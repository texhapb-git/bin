package main

import (
	"fmt"
	"log"

	"bin/api"
	"bin/file"
	"bin/storage"
)

func main() {
	// Создаем зависимости
	fileService := file.NewFileManager()
	storageService := storage.NewJSONStorage(fileService, "data/bins.json")
	binService := api.NewBinAPI(storageService)

	// Демонстрация работы с bin'ами
	fmt.Println("Создаем новый bin...")
	bin, err := binService.CreateBin("Мой первый bin", false)
	if err != nil {
		log.Fatalf("Ошибка создания bin: %v", err)
	}
	fmt.Printf("Создан bin: %+v\n", bin)

	fmt.Println("\nПолучаем все bins...")
	binList, err := binService.GetAllBins()
	if err != nil {
		log.Fatalf("Ошибка получения bins: %v", err)
	}
	fmt.Printf("Найдено bins: %d\n", len(binList.Bins))

	fmt.Println("\nПолучаем конкретный bin...")
	retrievedBin, err := binService.GetBin(bin.ID)
	if err != nil {
		log.Fatalf("Ошибка получения bin: %v", err)
	}
	fmt.Printf("Получен bin: %+v\n", retrievedBin)
}

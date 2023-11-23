package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Database представляет структуру данных для объекта из JSON
type Database struct {
	ID          string `json:"ID"`
	OVD         string `json:"OVD"`
	INSERT_DATE string `json:"INSERT_DATE"`
	NZ          string `json:"NZ"`
	IMEI        string `json:"IMEI"`
	NK          string `json:"NK"`
	DK          string `json:"DK"`
	DTL         string `json:"DTL"`
}

func main() {
	file, err := os.Open("mvswantedmt.json")
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	var db []Database // Изменение типа на срез структур Database

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&db); err != nil {
		fmt.Println("Ошибка при декодировании файла JSON:", err)
		return
	}

	// Теперь переменная db содержит данные из JSON в виде среза структур Database
	for _, item := range db {
		// Выводим содержимое каждого объекта базы данных
		fmt.Println("ID:", item.ID)
		fmt.Println("OVD:", item.OVD)
		fmt.Println("INSERT_DATE:", item.INSERT_DATE)
		fmt.Println("NZ:", item.NZ)
		fmt.Println("IMEI:", item.IMEI)
		fmt.Println("NK:", item.NK)
		fmt.Println("DK:", item.DK)
		fmt.Println("DTL:", item.DTL)
		fmt.Println()
	}
}

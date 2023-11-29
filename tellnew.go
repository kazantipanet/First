package main

import (
	"encoding/json"
	"fmt"
	"os"
)

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

	var db []Database

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&db); err != nil {
		fmt.Println("Ошибка при декодировании файла JSON:", err)
		return
	}

	// Запрос ввода IMEI у пользователя
	var searchIMEI string
	fmt.Print("Введите IMEI для поиска: ")
	fmt.Scanln(&searchIMEI)

	found := false

	// Поиск в базе данных
	for _, item := range db {
		if item.IMEI == searchIMEI {
			fmt.Println("Найденная запись:")
			fmt.Println("ID:", item.ID)
			fmt.Println("OVD:", item.OVD)
			fmt.Println("INSERT_DATE:", item.INSERT_DATE)
			fmt.Println("NZ:", item.NZ)
			fmt.Println("IMEI:", item.IMEI)
			fmt.Println("NK:", item.NK)
			fmt.Println("DK:", item.DK)
			fmt.Println("DTL:", item.DTL)

			fmt.Println()
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Запись с указанным IMEI в базе краденых телефонов не найдена")
	}
}

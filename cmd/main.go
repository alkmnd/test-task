package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"test-task/httpclient"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	baseURL := os.Getenv("BASE_URL")
	whatsappToken := os.Getenv("WHATSAPP_TOKEN")

	client := httpclient.NewClient(baseURL, whatsappToken)

	var currency, crm string
	fmt.Print("Введите валюту (например, RUB): ")
	_, err = fmt.Scanln(&currency)
	if err != nil {
		fmt.Println("Ошибка при вводе валюты:", err)
		return
	}

	fmt.Print("Введите CRM (например, lk): ")
	_, err = fmt.Scanln(&crm)
	if err != nil {
		fmt.Println("Ошибка при вводе CRM:", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	json, err := client.GetTariffs(ctx, currency, crm)
	if err != nil {
		fmt.Println("Ошибка при отправке запроса:", err)
		return
	}
	fmt.Println(json)

}

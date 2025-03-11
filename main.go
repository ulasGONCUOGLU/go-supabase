package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Log struct {
	ID      int    `json:"id"`
	Message string `json:"idsi"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(".env dosyası yüklenemedi: %v", err)
	}

	supabaseURL := os.Getenv("SUPABASE_URL")
	supabaseAnonKey := os.Getenv("SUPABASE_ANON_KEY")

	if supabaseURL == "" || supabaseAnonKey == "" {
		log.Fatal("SUPABASE_URL veya SUPABASE_ANON_KEY tanımlı değil")
	}

	tableName := "logs"
	requestURL := fmt.Sprintf("%s/rest/v1/%s", supabaseURL, tableName)

	client := &http.Client{}
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		log.Fatalf("İstek oluşturulamadı: %v", err)
	}

	req.Header.Set("apikey", supabaseAnonKey)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", supabaseAnonKey))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("İstek gönderilemedi: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		log.Fatalf("Hata: %s", string(bodyBytes))
	}

	var logs []Log
	if err := json.NewDecoder(resp.Body).Decode(&logs); err != nil {
		log.Fatalf("JSON çözümlenemedi: %v", err)
	}

	// 3) Ekrana sonuçları bas
	fmt.Println("Tüm loglar:")
	for _, logEntry := range logs {
		fmt.Printf("ID: %d, Mesaj: %s\n", logEntry.ID, logEntry.Message)
	}

	// 4) Ekrana id 1 olanı bas
	fmt.Println("\nID'si 1 olan log:")
	for _, logEntry := range logs {
		if logEntry.ID == 1 {
			fmt.Printf("ID: %d, Mesaj: %s\n", logEntry.ID, logEntry.Message)
			break
		}
	}

	// 5) Yeni veri ekleme isteği at (içeriği boş olsun, sadece insert isteği atalım)
	newLog := Log{Message: "Yeni log mesajı"}
	newLogJSON, err := json.Marshal(newLog)
	if err != nil {
		log.Fatalf("JSON oluşturulamadı: %v", err)
	}

	req, err = http.NewRequest("POST", requestURL, ioutil.NopCloser(bytes.NewReader(newLogJSON)))
	if err != nil {
		log.Fatalf("İstek oluşturulamadı: %v", err)
	}

	req.Header.Set("apikey", supabaseAnonKey)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", supabaseAnonKey))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Prefer", "return=representation")

	resp, err = client.Do(req)
	if err != nil {
		log.Fatalf("İstek gönderilemedi: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		log.Fatalf("Hata: %s", string(bodyBytes))
	}

	var insertedLogs []Log
	if err := json.NewDecoder(resp.Body).Decode(&insertedLogs); err != nil {
		log.Fatalf("JSON çözümlenemedi: %v", err)
	}

	fmt.Println("\nEklenen log:")
	for _, logEntry := range insertedLogs {
		fmt.Printf("ID: %d, Mesaj: %s\n", logEntry.ID, logEntry.Message)
	}
}

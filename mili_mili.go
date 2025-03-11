/*
package mili_mili

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"myproject/models"
	"net/http"
	"os"
)
a
func UpdateCounter(w http.ResponseWriter, r *http.Request) {
	// Gelen JSON verisini oku
	var counter models.Counter
	body, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(body, &counter)
	if err != nil {
		http.Error(w, "Geçersiz JSON verisi!", http.StatusBadRequest)
		return
	}

	// Supabase API URL'sini oluştur
	supabaseURL := os.Getenv("SUPABASE_URL")
	if supabaseURL == "" {
		http.Error(w, "Supabase URL çevresel değişkeni boş!", http.StatusInternalServerError)
		return
	}
	url := fmt.Sprintf("%s/rest/v1/counter?id=eq.%d", supabaseURL, counter.ID)

	// Supabase'den mevcut 'value' değerini al
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("apikey", os.Getenv("SUPABASE_KEY"))
	req.Header.Set("Authorization", "Bearer "+os.Getenv("SUPABASE_KEY"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		http.Error(w, "Supabase ile iletişim hatası", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var result []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil || len(result) == 0 {
		http.Error(w, "ID'ye sahip veri bulunamadı", http.StatusNotFound)
		return
	}

	// Mevcut value'yu al ve 1 artır
	currentValue := result[0]["value"].(float64)
	newValue := int(currentValue) + 1

	// Yeni value değerini Supabase'e göndermek için JSON verisi hazırla
	updateData := map[string]int{"value": newValue}
	updateJson, _ := json.Marshal(updateData)

	// Supabase'e PATCH isteği gönder
	req, _ = http.NewRequest("PATCH", url, bytes.NewBuffer(updateJson))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("apikey", os.Getenv("SUPABASE_KEY"))
	req.Header.Set("Authorization", "Bearer "+os.Getenv("SUPABASE_KEY"))
	req.Header.Set("Prefer", "return=minimal")

	resp, err = client.Do(req)
	if err != nil || resp.StatusCode != http.StatusNoContent {
		http.Error(w, "Veri güncelleme başarısız", http.StatusInternalServerError)
		return
	}

	// Başarılı yanıt döndür
	w.WriteHeader(http.StatusOK)
	w.Write([]byte({"message": "Veri başarıyla güncellendi ve değeri 1 arttırıldı!"}))
}*/;
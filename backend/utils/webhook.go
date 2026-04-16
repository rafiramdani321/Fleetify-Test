package utils

import (
	"bytes"
	"encoding/json"
	"fleetify-backend/models"
	"fmt"
	"net/http"
	"os"
	"time"
)

func SendInvoiceWebhook(invoice models.Invoice) {
	go func (data models.Invoice) {
		webhookURL := os.Getenv("WEB_HOOK")
		if webhookURL == "" {
			fmt.Println("Webhook URL not set in env")
			return
		}

		jsonData, _ := json.Marshal(data)

		client := &http.Client{Timeout: 10 * time.Second}
		resp, err := client.Post(webhookURL, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Printf("Webhook Error: %v\n", err)
			return
		}
		defer resp.Body.Close()

		fmt.Printf("Webhook Send! Status: %s\n", resp.Status)
	}(invoice)
}
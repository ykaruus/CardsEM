package storage

import (
	"bytes"
	"encoding/json"
	"kairusService/internal/domain/entities"
	"kairusService/internal/domain/repository"
	"net/http"
)

type DiscordWebhookStorage struct {
	debug_url    string
	alert_url    string
	notification string
}

func NewDiscordWebhookStorage(
	debug_url string,
	alert_url string,
	notification string,
) *DiscordWebhookStorage {
	return &DiscordWebhookStorage{
		debug_url:    debug_url,
		alert_url:    alert_url,
		notification: notification,
	}
}

var _ repository.Discord_Webhook_Repository = (*DiscordWebhookStorage)(nil)

func (w *DiscordWebhookStorage) SendAlert(embed entities.Embed) (int, error) {
	body, err := json.Marshal(map[string]interface{}{
		"content": "@everyone",
		"embeds": []map[string]interface{}{
			{
				"title":       embed.Title,
				"description": embed.Description,
				"color":       3447003,
			},
		},
	})

	if err != nil {
		return 0, err
	}

	payload := bytes.NewBuffer(body)

	resp, err := http.Post(w.alert_url, "application/json", payload)

	if err != nil {
		return 0, err
	}

	return resp.StatusCode, nil

}

// func (w *DiscordWebhookStorage) SendError(interface{}) error {

// }

// func (w *DiscordWebhookStorage) SendNotification(interface{}) error {

// }

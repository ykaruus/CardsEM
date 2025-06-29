package services

import (
	"kairusService/internal/domain/entities"
	"kairusService/internal/storage"
)

type DiscordWebhookService struct {
	webhook_storage *storage.DiscordWebhookStorage
}

func NewDiscordWebhookService(webhook_storage *storage.DiscordWebhookStorage) *DiscordWebhookService {
	return &DiscordWebhookService{
		webhook_storage: webhook_storage,
	}
}

func (service *DiscordWebhookService) SendAlert(embed entities.Embed) (int, error) {
	responseCode, err := service.webhook_storage.SendAlert(embed)

	if err != nil {
		return 0, err
	}

	return responseCode, nil
}

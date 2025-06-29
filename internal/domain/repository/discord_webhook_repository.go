package repository

import "kairusService/internal/domain/entities"

type Discord_Webhook_Repository interface {
	SendAlert(entities.Embed) (int, error)
	// SendError(interface{}) error
	// SendNotification(interface{}) error
}

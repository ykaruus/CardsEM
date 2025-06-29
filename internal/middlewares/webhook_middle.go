package middlewares

import (
	"github.com/gin-gonic/gin"
)

type WebhookMiddle struct {
}

func NewWebhookMiddle() *WebhookMiddle {
	return &WebhookMiddle{}
}

func (m *WebhookMiddle) Webhook_handler(apiKeySecret string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var api_key string = ctx.GetHeader("X-API-Key")

		if api_key == "" || api_key != apiKeySecret {
			ctx.AbortWithStatusJSON(401, gin.H{
				"message": "Acesso não autorizado, é necessario uma chave de api...",
			})
			return
		}

		ctx.Next()

	}
}

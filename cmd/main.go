package main

import (
	"fmt"
	"kairusService/internal/container"
	"kairusService/internal/domain/entities"
	"kairusService/internal/infra"
	"kairusService/internal/middlewares"
	"kairusService/internal/services"
	"kairusService/internal/storage"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()

	err := godotenv.Load()

	if err != nil {
		fmt.Println("Ocorreu um erro ao carregar um arquivo de ambiente: ", err.Error())
		return
	}

	dbUri := os.Getenv("DB_URL")
	dbName := os.Getenv("DB_NAME")
	dbUsers := os.Getenv("DB_USERS")
	secretkey := os.Getenv("SECRET_KEY")
	apiSecretKey := os.Getenv("BOT_DISCORD_API_KEY")
	webhook_url := os.Getenv("ALERT_WEBHOOK")

	client, err := storage.ConnectMongo(dbUri)

	if err != nil {
		fmt.Println("Erro ao conectar com a database..", err)
		return
	}

	fmt.Println(dbUri, dbName, dbUsers, secretkey, apiSecretKey)

	webhook_storage := storage.NewDiscordWebhookStorage("", webhook_url, "")

	webhook_service := services.NewDiscordWebhookService(webhook_storage)

	webhook_service.SendAlert(entities.Embed{
		Title:       "--------",
		Description: "--------",
	})

	users_collection := storage.LoadCollection(client, dbName, dbUsers)
	userStorage := storage.NewMongoRepository(users_collection)

	container := container.NewContainer(secretkey, userStorage)

	userController, authController := infra.InjectUsersServices(container.UserService, container.ApiResponse, container.AuthService)
	AuthMiddleware := middlewares.NewAuthMiddleware(container.TokenService, container.ApiResponse)

	users := router.Group("/api/v1/admin")
	{
		users.Use(AuthMiddleware.CheckAdmin())
		users.POST("/users", userController.Create)
		users.GET("/users/:id", userController.GetUserFromId)
		users.GET("/users", userController.GetUsers)
		users.GET("/users/query", userController.GetUserFromName)
		users.PUT("/users/:id", userController.UpdateUser)
	}

	auth := router.Group("/api/v1/auth")
	{
		auth.POST("/login", authController.Login)
		auth.GET("/decript", authController.DecriptToken)
	}
	fmt.Println("Rodando o servidor na rota http://localhost:5000")
	router.Run("localhost:5000")

}

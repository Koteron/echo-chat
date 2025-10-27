package main

import (
	"ChatService/client"
	"ChatService/internal/controller"
	"ChatService/internal/db"
	"ChatService/internal/repository"
	"ChatService/internal/service"
	"ChatService/internal/transaction"

	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := db.Init(os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	if err != nil {
		log.Fatal("Failed to connect to db: ", err)
		return
	}

	chatRepository := repository.NewChatRepo(db)
	chatUserRepository := repository.NewChatUserRepo(db)
	unitOfWork := transaction.NewUnitOfWork(db)

	userServiceGRPC, err := client.NewUserServiceClient(os.Getenv("USER_SERVICE_GRPC_HOST"))

	if err != nil {
		log.Fatal("Failed to connect to user service grpc: ", err)
		return
	}

	chatUserService := service.NewChatUserService(chatUserRepository)
	chatService := service.NewChatService(chatRepository, chatUserService, userServiceGRPC, unitOfWork)
	chatController := controller.NewChatController(chatService)

	r := gin.Default()

	r.GET("/contacts/:id", chatController.GetAllChatsByUserId)
	r.GET("/search", chatController.SearchChats)

	r.POST("/create", chatController.CreateChat)

	r.Run(":8080")
}

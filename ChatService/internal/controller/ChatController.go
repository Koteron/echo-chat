package controller

import (
	"ChatService/internal/dto"
	"ChatService/internal/service"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ChatController struct {
	service *service.ChatService
}

func NewChatController(s *service.ChatService) *ChatController {
	return &ChatController{service: s}
}

func (c *ChatController) GetAllChatsByUserId(context *gin.Context) {
	userId, err := uuid.Parse(context.Param("id"))

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	chats, err := c.service.GetAllChatsByUserId(context, userId)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	context.JSON(http.StatusOK, chats)
}

func (c *ChatController) CreateChat(gitCtx *gin.Context) {
	var req dto.ChatCreationDto

	if err := gitCtx.ShouldBindJSON(&req); err != nil {
		gitCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := gitCtx.Request.Context()

	chatDisplayDto, err := c.service.CreateChat(ctx, &req)
	if err != nil {
		gitCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	gitCtx.JSON(http.StatusCreated, chatDisplayDto)
}

func (c *ChatController) SearchChats(ginCtx *gin.Context) {
	var req dto.SearchChatsRequest
	if err := ginCtx.ShouldBindQuery(&req); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Page < 1 {
		req.Page = 1
	}
	if req.Limit <= 0 || req.Limit > 100 {
		req.Limit = 20
	}

	ctx := ginCtx.Request.Context()
	result, err := c.service.SearchChats(ctx, req)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ginCtx.JSON(http.StatusOK, result)
}

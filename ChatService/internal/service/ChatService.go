package service

import (
	"context"

	"ChatService/client"
	"ChatService/internal/dto"
	"ChatService/internal/entity"
	"ChatService/internal/entity/util"
	"ChatService/internal/repository"
	"ChatService/internal/transaction"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ChatService struct {
	repo            *repository.ChatRepo
	chatUserService *ChatUserService
	userServiceGRPC *client.UserServiceGRPC
	unitOfWork      *transaction.UnitOfWork
}

func NewChatService(repo *repository.ChatRepo,
	chatUserService *ChatUserService, userServiceGRPC *client.UserServiceGRPC,
	unitOfWork *transaction.UnitOfWork) *ChatService {
	return &ChatService{repo, chatUserService, userServiceGRPC, unitOfWork}
}

func (s *ChatService) GetAllChatsByUserId(ctx context.Context, id uuid.UUID) ([]dto.ChatDisplayDto, error) {
	directUserIds, err := s.chatUserService.GetDirectUserIdsForUser(ctx, id)

	if err != nil {
		return nil, err
	}

	directUserIdsString := make([]string, len(directUserIds))
	for i, id := range directUserIds {
		directUserIdsString[i] = id.String()
	}

	directUserNames, err := s.userServiceGRPC.GetDisplayNames(directUserIdsString)

	if err != nil {
		return nil, err
	}

	chats, err := s.repo.GetAllChatsByUserId(ctx, nil, id)

	if err != nil {
		return nil, err
	}

	chatDtos := dto.ToChatDisplayDtos(chats)

	for i, chatDto := range chatDtos {
		if chatDto.Type == util.Direct {
			chatDto.Title = &directUserNames[i] // Is there an order guarantee???
		}
	}

	return chatDtos, nil
}

func (s *ChatService) CreateChat(ctx context.Context, chatCreationDto *dto.ChatCreationDto) (dto.ChatDisplayDto, error) {
	var chat *entity.Chat
	err := s.unitOfWork.Do(ctx, func(tx *gorm.DB) error {

		chat, err := s.repo.Create(ctx, tx, &entity.Chat{
			Id:    uuid.New(),
			Type:  chatCreationDto.Type,
			Title: chatCreationDto.Title,
		})
		if err != nil {
			return err
		}

		if err := s.chatUserService.AddMembers(ctx, tx, chatCreationDto.UserIds, chat.Id); err != nil {
			return err
		}

		return nil
	})
	return dto.ToChatDisplayDto(chat), err
}

func (s *ChatService) SearchChats(ctx context.Context, req dto.SearchChatsRequest) (
	*dto.PaginatedResponse[dto.ChatDisplayDto], error) {
	chats, total, err := s.repo.SearchChats(ctx, nil, req.Query, req.Page, req.Limit)
	if err != nil {
		return nil, err
	}

	totalPages := int((total + int64(req.Limit) - 1) / int64(req.Limit))
	return &dto.PaginatedResponse[dto.ChatDisplayDto]{
		Data:       dto.ToChatDisplayDtos(chats),
		Total:      total,
		Page:       req.Page,
		Limit:      req.Limit,
		TotalPages: totalPages,
	}, nil
}

package service

import (
	"ChatService/internal/entity"
	"ChatService/internal/repository"

	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ChatUserService struct {
	repo *repository.ChatUserRepo
}

func NewChatUserService(repo *repository.ChatUserRepo) *ChatUserService {
	return &ChatUserService{repo}
}

func (s *ChatUserService) GetDirectUserIdsForUser(ctx context.Context, userId uuid.UUID) ([]uuid.UUID, error) {
	return s.repo.GetDirectUserIdsForUser(ctx, nil, userId)
}

func (s *ChatUserService) LeaveChat(ctx context.Context, userId uuid.UUID, chatId uuid.UUID) error {
	return s.repo.SoftDeleteById(ctx, nil, userId, chatId)
}

func (s *ChatUserService) AddMembers(ctx context.Context, tx *gorm.DB, userIds []uuid.UUID, chatId uuid.UUID) error {
	chatUsers := make([]entity.ChatUser, len(userIds))
	for i, userId := range userIds {
		chatUsers[i] = entity.ChatUser{
			ChatId: chatId,
			UserId: userId,
		}
	}
	return s.repo.AddMembers(ctx, tx, chatUsers)
}

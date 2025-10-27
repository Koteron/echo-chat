package repository

import (
	"ChatService/internal/entity"
	"ChatService/internal/entity/util"

	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ChatRepo struct {
	db *gorm.DB
}

func NewChatRepo(db *gorm.DB) *ChatRepo {
	return &ChatRepo{db}
}

func (r *ChatRepo) getDB(ctx context.Context, tx *gorm.DB) *gorm.DB {
	if tx != nil {
		return tx.WithContext(ctx)
	}
	return r.db.WithContext(ctx)
}

func (r *ChatRepo) Create(ctx context.Context, tx *gorm.DB, chat *entity.Chat) (*entity.Chat, error) {
	db := r.getDB(ctx, tx)

	return chat, db.Create(chat).Error
}

func (r *ChatRepo) GetById(ctx context.Context, tx *gorm.DB, id uuid.UUID) (*entity.Chat, error) {
	var chat entity.Chat
	err := r.getDB(ctx, tx).First(&chat, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &chat, nil
}

func (r *ChatRepo) DeleteById(ctx context.Context, tx *gorm.DB, id uuid.UUID) error {
	return r.getDB(ctx, tx).Delete(&entity.Chat{}, "id = ?", id).Error
}

func (r *ChatRepo) GetAllChatsByUserId(ctx context.Context, tx *gorm.DB, id uuid.UUID) ([]entity.Chat, error) {
	var chats []entity.Chat
	db := r.getDB(ctx, tx)

	result := db.
		Preload("Users").
		Joins("JOIN chat_users cu ON cu.chat_id = chats.id").
		Where("cu.user_id = ?", id).
		Where("cu.is_active = true").
		Find(&chats)

	if result.Error != nil {
		return nil, result.Error
	}

	return chats, nil
}

func (r *ChatRepo) SearchChats(ctx context.Context, tx *gorm.DB, query string, page, limit int) ([]entity.Chat, int64, error) {
	var chats []entity.Chat
	var total int64

	db := r.getDB(ctx, tx)

	if query != "" {
		db = db.Where("title ILIKE ? AND type = ?", "%"+query+"%", util.PublicGroup)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	if err := db.Order("created_at DESC").Limit(limit).Offset(offset).Find(&chats).Error; err != nil {
		return nil, 0, err
	}

	return chats, total, nil
}

package repository

import (
	"ChatService/internal/entity"
	"ChatService/internal/entity/util"
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ChatUserRepo struct {
	db *gorm.DB
}

func NewChatUserRepo(db *gorm.DB) *ChatUserRepo {
	return &ChatUserRepo{db}
}

func (r *ChatUserRepo) getDB(ctx context.Context, tx *gorm.DB) *gorm.DB {
	if tx != nil {
		return tx.WithContext(ctx)
	}
	return r.db.WithContext(ctx)
}

func (r *ChatUserRepo) SaveChatUser(ctx context.Context, tx *gorm.DB, chatUser *entity.ChatUser) error {
	db := r.getDB(ctx, tx)
	result := db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(chatUser)
	return result.Error
}

func (r *ChatUserRepo) SoftDeleteById(ctx context.Context, tx *gorm.DB, userId, chatId uuid.UUID) error {
	db := r.getDB(ctx, tx)

	var chatUser entity.ChatUser
	if err := db.First(&chatUser, "chat_id = ? AND user_id = ?", chatId, userId).Error; err != nil {
		return err
	}

	chatUser.IsActive = false
	now := time.Now()
	chatUser.LeftAt = &now

	return r.SaveChatUser(ctx, tx, &chatUser)
}

func (r *ChatUserRepo) GetDirectUserIdsForUser(ctx context.Context, tx *gorm.DB, userId uuid.UUID) ([]uuid.UUID, error) {
	db := r.getDB(ctx, tx)

	var directUserIDs []uuid.UUID
	err := db.Table("chat_users AS cu1").
		Select("DISTINCT cu2.user_id").
		Joins("JOIN chats c ON cu1.chat_id = c.id").
		Joins("JOIN chat_users cu2 ON cu1.chat_id = cu2.chat_id").
		Where(`
            cu1.user_id = ? 
            AND c.type = ? 
            AND cu2.user_id <> cu1.user_id 
            AND cu1.is_active = true
        `, userId, util.Direct).
		Pluck("cu2.user_id", &directUserIDs).Error

	if err != nil {
		return nil, err
	}
	return directUserIDs, nil
}

func (r *ChatUserRepo) AddMembers(ctx context.Context, tx *gorm.DB, members []entity.ChatUser) error {
	db := tx
	if db == nil {
		db = r.db.WithContext(ctx)
	}
	return db.CreateInBatches(members, 100).Error
}

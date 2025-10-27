package entity

import (
	"ChatService/internal/entity/util"
	"time"

	"github.com/google/uuid"
)

type Chat struct {
	Id          uuid.UUID     `gorm:"primaryKey"`
	Type        util.ChatType `gorm:"not null"`
	Title       *string
	Description *string
	Users       []ChatUser `gorm:"foreignKey:ChatID"`
	UpdatedAt   time.Time  `gorm:"autoCreateTime"`
	CreatedAt   time.Time  `gorm:"autoCreateTime"`
}

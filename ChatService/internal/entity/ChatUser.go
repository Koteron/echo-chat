package entity

import (
	"ChatService/internal/entity/util"
	"time"

	"github.com/google/uuid"
)

type ChatUser struct {
	ChatId      uuid.UUID  `gorm:"primaryKey"`
	UserId      uuid.UUID  `gorm:"primaryKey"`
	Role        util.Role  `gorm:"default:MEMBER"`
	IsBanned    bool       `gorm:"default:false"`
	BannedAt    *time.Time `gorm:"default:null"`
	BannedUntil *time.Time `gorm:"default:null"`
	PromotedBy  *uuid.UUID `gorm:"default:null"`
	PromotedAt  *time.Time `gorm:"default:null"`
	JoinedAt    time.Time  `gorm:"autoCreateTime"`
	LeftAt      *time.Time `gorm:"default:null"`
	IsActive    bool       `gorm:"default:true"`
	Chat        Chat       `gorm:"foreignKey:ChatID"`
}

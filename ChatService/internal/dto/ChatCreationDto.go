package dto

import (
	"ChatService/internal/entity/util"

	"github.com/google/uuid"
)

type ChatCreationDto struct {
	UserIds []uuid.UUID
	Title   *string
	Type    util.ChatType
}

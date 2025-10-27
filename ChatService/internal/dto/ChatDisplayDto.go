package dto

import (
	"ChatService/internal/entity"
	"ChatService/internal/entity/util"

	"github.com/google/uuid"
)

type ChatDisplayDto struct {
	Id    uuid.UUID
	Type  util.ChatType
	Title *string
}

func ToChatDisplayDto(chat *entity.Chat) ChatDisplayDto {
	title := ""
	if chat.Title != nil {
		title = *chat.Title
	}
	return ChatDisplayDto{
		Id:    chat.Id,
		Title: &title,
		Type:  chat.Type,
	}
}

func ToChatDisplayDtos(chats []entity.Chat) []ChatDisplayDto {
	dtos := make([]ChatDisplayDto, len(chats))
	for i, c := range chats {
		dtos[i] = ToChatDisplayDto(&c)
	}
	return dtos
}

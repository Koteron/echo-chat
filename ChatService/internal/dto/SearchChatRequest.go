package dto

type SearchChatsRequest struct {
    Query string `form:"query"`
    Page  int    `form:"page"`
    Limit int    `form:"limit"`
}
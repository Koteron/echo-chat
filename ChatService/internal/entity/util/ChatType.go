package util

type ChatType string

const (
	Direct       ChatType = "DIRECT"
	PrivateGroup ChatType = "PRIVATE_GROUP"
	PublicGroup  ChatType = "PUBLIC_GROUP"
)

package db

import (
	"fmt"

	"ChatService/internal/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Init(host string, user string, password string, name string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		host, user, password, name,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: false,
			NameReplacer:  nil,
		},
	})

	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&entity.Chat{}, &entity.ChatUser{})

	if err != nil {
		return nil, err
	}

	return db, nil
}

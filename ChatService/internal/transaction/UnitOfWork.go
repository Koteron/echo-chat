package transaction

import (
	"context"

	"gorm.io/gorm"
)

type UnitOfWork struct {
    db *gorm.DB
}

func NewUnitOfWork(db *gorm.DB) *UnitOfWork {
    return &UnitOfWork{db}
}

func (u *UnitOfWork) Do(ctx context.Context, fn func(tx *gorm.DB) error) error {
    tx := u.db.WithContext(ctx).Begin()
    defer func() {
        if r := recover(); r != nil {
            tx.Rollback()
            panic(r)
        }
    }()

    if err := fn(tx); err != nil {
        tx.Rollback()
        return err
    }
    return tx.Commit().Error
}
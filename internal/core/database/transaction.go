package database

import (
	"context"

	"gorm.io/gorm"
)

type TxManager interface {
	WithTx(ctx context.Context, fn func(ctx context.Context) error) error
}

type GormTxManager struct {
	db *gorm.DB
}

func NewTxManager(db *gorm.DB) *GormTxManager {
	return &GormTxManager{db: db}
}

type txKeyType struct{}

var txKey = txKeyType{}

func (tm *GormTxManager) WithTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return tm.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		txCtx := context.WithValue(ctx, txKey, tx)
		return fn(txCtx)
	})
}

func GetTxFromContext(ctx context.Context, defaultDB *gorm.DB) *gorm.DB {
	if tx, ok := ctx.Value(txKey).(*gorm.DB); ok {
		return tx
	}
	return defaultDB
}

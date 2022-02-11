package component

import (
	"gorm.io/gorm"
)

type AppContext interface {
	GetDbConnection() *gorm.DB
	GetSecretKey() string
	GetBcryptCost() int
}

type appContext struct {
	db         *gorm.DB
	secretKey  string
	bcryptCost int
}

func NewAppContext(db *gorm.DB, secretKey string, bcryptCost int) *appContext {
	return &appContext{db: db, secretKey: secretKey, bcryptCost: bcryptCost}
}

func (ctx *appContext) GetDbConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appContext) GetSecretKey() string {
	return ctx.secretKey
}

func (ctx *appContext) GetBcryptCost() int {
	return ctx.bcryptCost
}

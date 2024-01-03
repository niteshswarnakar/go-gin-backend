package service

import (
	"github.com/render-examples/go-gin-web-server/database"
	"github.com/render-examples/go-gin-web-server/infrastructure"
	"gorm.io/gorm"
)

type Service struct {
	db  *database.DataBase
	env infrastructure.Env
}

func (s Service) GetDB() *gorm.DB {
	db := s.db.Conn.Db
	return db
}

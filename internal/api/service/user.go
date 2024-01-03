package service

import (
	"github.com/render-examples/go-gin-web-server/database"
	"github.com/render-examples/go-gin-web-server/infrastructure"
	"github.com/render-examples/go-gin-web-server/internal/domain"
	"github.com/render-examples/go-gin-web-server/internal/view"
)

type UserService interface {
	Create(view view.UserCreateView) (*domain.User, error)
}

type userService struct {
	Service
}

func NewUserService(db *database.DataBase, env *infrastructure.Env) UserService {
	return &userService{
		Service{
			db:  db,
			env: *env,
		},
	}
}

func (u *userService) Create(view view.UserCreateView) (*domain.User, error) {
	db := u.GetDB()
	user := &domain.User{
		Name:  view.Name,
		Email: view.Email,
	}
	if err := db.Create(&user).Error; err != nil {
		return nil, err
	}
	return user,nil

}

package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/render-examples/go-gin-web-server/internal/api/service"
	"github.com/render-examples/go-gin-web-server/internal/response"
	"github.com/render-examples/go-gin-web-server/internal/view"
)

type userHandler struct {
	service service.UserService
}

type UserHandler interface {
	Create(ctx *gin.Context)
}

func NewUserHandler(service service.UserService) UserHandler {
	return &userHandler{
		service: service,
	}
}

func (u *userHandler) Create(ctx *gin.Context) {
	var userView view.UserCreateView
	err := ctx.ShouldBindJSON(&userView)
	if err != nil {
		response.BadRequestResponse(ctx, "Bad request")
		return
	}
	user, err := u.service.Create(userView)
	if err != nil {
		response.ServerErrorResponse(ctx, err)
		return
	}
	response.SuccessResponse(ctx, user)

}

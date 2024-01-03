package route

import (
	"github.com/gin-gonic/gin"
	"github.com/render-examples/go-gin-web-server/database"
	"github.com/render-examples/go-gin-web-server/infrastructure"
	handler2 "github.com/render-examples/go-gin-web-server/internal/api/handler"
	service2 "github.com/render-examples/go-gin-web-server/internal/api/service"
)

func UserRoutes(group *gin.RouterGroup, db *database.DataBase, env *infrastructure.Env) {
	service := service2.NewUserService(db, env)
	handler := handler2.NewUserHandler(service)
	group.POST("/register", handler.Create)
}

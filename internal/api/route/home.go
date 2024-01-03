package route

import (
	"github.com/gin-gonic/gin"
	"github.com/render-examples/go-gin-web-server/internal/api/handler"
)

func HomeRoutes(group *gin.RouterGroup) {
	group.GET("/", handler.HomeView)
}

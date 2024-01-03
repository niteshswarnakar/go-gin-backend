package route

import (
	"github.com/gin-gonic/gin"
	"github.com/render-examples/go-gin-web-server/bootstrap"
)

func Setup(app *bootstrap.Application, r *gin.Engine) {
	group := r.Group("/api")
	HomeRoutes(group)
	UserRoutes(group, app.Database, app.Env)
}

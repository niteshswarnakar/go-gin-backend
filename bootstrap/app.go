package bootstrap

import (
	"github.com/render-examples/go-gin-web-server/database"
	"github.com/render-examples/go-gin-web-server/infrastructure"
)

type Application struct {
	Env      *infrastructure.Env
	Database *database.DataBase
}

func App() Application {
	app := &Application{}
	env := infrastructure.NewEnv()
	dbConn := infrastructure.NewDbConnection(env)
	app.Database = database.NewDataBase(dbConn)
	app.Env = &env
	return *app
}

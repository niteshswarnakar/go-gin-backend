package main

import (
	"github.com/render-examples/go-gin-web-server/infrastructure"
	"github.com/render-examples/go-gin-web-server/internal/domain"
	"log"
	"os"
)

func main() {
	env := infrastructure.NewEnv()
	db := infrastructure.NewDbConnection(env)
	tx := db.Db.Begin()
	err := tx.Migrator().DropTable(domain.AllModels...)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		os.Exit(1)
	}
	err = tx.AutoMigrate(domain.AllModels...)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		os.Exit(1)
	}
	tx.Commit()

}

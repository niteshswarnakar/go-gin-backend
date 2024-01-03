package main

import (
	"fmt"
	"github.com/render-examples/go-gin-web-server/bootstrap"
	"github.com/render-examples/go-gin-web-server/internal/api/route"
	"log"
	"runtime"

	"github.com/gin-gonic/gin"
)

func main() {
	//ConfigRuntime()
	app := bootstrap.App()
	env := app.Env
	router := gin.New()
	route.Setup(&app, router)
	if err := router.Run(":" + env.ServerPort); err != nil {
		log.Panicf("Can't spin up the server: %s", err)
	}
}

// ConfigRuntime sets the number of operating system threads.
func ConfigRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
}

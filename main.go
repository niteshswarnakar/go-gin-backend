package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"
)

func main() {
	// ConfigRuntime()
	StartGin()
}

// ConfigRuntime sets the number of operating system threads.
func ConfigRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
}

// StartGin starts gin web server with setting router.
func StartGin() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.GET("/", HomeView)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	fmt.Println("server started at ", port)

	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}

type Route struct {
	Url    string
	Method string
}

func HomeView(c *gin.Context) {
	var routes []*Route

	homeRoute := &Route{
		Method: "GET",
		Url:    "/",
	}

	createBookRoute := &Route{
		Method: "POST",
		Url:    "/book",
	}

	getBookRoute := &Route{
		Method: "GET",
		Url:    "/book/:id",
	}

	listBookRoute := &Route{
		Method: "GET",
		Url:    "/books",
	}

	deleteBookRoute := &Route{
		Method: "DELETE",
		Url:    "/book/:id",
	}

	createConfession := &Route{
		Method: "POST",
		Url:    "/confession",
	}

	routes = append(routes, homeRoute, createBookRoute, getBookRoute, listBookRoute, deleteBookRoute, createConfession)

	c.IndentedJSON(http.StatusAccepted, routes)
}

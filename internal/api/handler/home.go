package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

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

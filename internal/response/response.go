package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	ErrorMessage string      `json:"message,omitempty"`
	Data         interface{} `json:"data,omitempty"`
	Error        bool        `json:"error"`
}

func SuccessResponse(ctx *gin.Context, data interface{}) {
	ctx.IndentedJSON(http.StatusOK, Response{
		Data:  data,
		Error: false,
	})
}

func ServerErrorResponse(ctx *gin.Context, err error) {
	ctx.IndentedJSON(http.StatusInternalServerError, Response{
		ErrorMessage: "Something went wrong",
		Error:        true,
	})
}

func BadRequestResponse(ctx *gin.Context, message string) {
	ctx.IndentedJSON(http.StatusBadRequest, Response{
		ErrorMessage: message,
		Error:        true,
	})
}

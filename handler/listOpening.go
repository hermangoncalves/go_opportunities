package handler

import "github.com/gin-gonic/gin"

func ListOpeningHandler(ctx *gin.Context) {
	sendSuccess(ctx, "list-opening", map[string]interface{}{
		"name": "Herman",
		"age": 23,
	})
}

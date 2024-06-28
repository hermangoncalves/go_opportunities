package handler

import "github.com/gin-gonic/gin"

func CreateOpeningHandler(ctx *gin.Context) {
	sendSuccess(ctx, "create-opening", map[string]interface{}{
		"name": "Herman",
		"age": 23,
	})
}

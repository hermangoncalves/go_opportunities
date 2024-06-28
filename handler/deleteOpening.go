package handler

import "github.com/gin-gonic/gin"

func DeleteOpeningHandler(ctx *gin.Context) {
	sendSuccess(ctx, "delete-opening", map[string]interface{}{
		"name": "Herman",
		"age": 23,
	})
}

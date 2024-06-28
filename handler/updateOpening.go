package handler

import "github.com/gin-gonic/gin"

func UpdateOpeningHandler(ctx *gin.Context) {
	sendSuccess(ctx, "update-opening", map[string]interface{}{
		"name": "Herman",
		"age": 23,
	})
}

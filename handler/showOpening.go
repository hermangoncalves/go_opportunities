package handler

import "github.com/gin-gonic/gin"

func ShowOpeningHandler(ctx *gin.Context) {
	opening := map[string]interface{} {
		
	}
	sendSuccess(ctx, "show-opening", opening)
}
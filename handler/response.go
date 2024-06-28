package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Operation from handler: %s successfull", op),
		"data": data,
	})
}

package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hermangoncalves/go_opportunities/schemas"
)

func ListOpeningHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}
	// if err := db.Find
	sendSuccess(ctx, "list-opening", openings)
}

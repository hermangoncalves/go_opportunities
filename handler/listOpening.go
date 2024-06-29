package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hermangoncalves/go_opportunities/schemas"
)

func ListOpeningHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}
	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}
	sendSuccess(ctx, "list-opening", openings)
}

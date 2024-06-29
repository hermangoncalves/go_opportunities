package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hermangoncalves/go_opportunities/schemas"
)

func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "string").Error())
		return
	}

	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		fmt.Printf("ShowOpening error: %v", err.Error)
		sendError(ctx, http.StatusBadRequest, "opening not found")
		return
	}

	sendSuccess(ctx, "show-opening", opening)
}

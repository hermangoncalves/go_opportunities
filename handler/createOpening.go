package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hermangoncalves/go_opportunities/schemas"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}
	
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}


	newOpening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&newOpening).Error; err != nil {
		logger.Errorf("error creating new opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error on creating new opening")
	}

	sendSuccess(ctx, "create-opening", newOpening)
}

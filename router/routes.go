package router

import (

	"github.com/gin-gonic/gin"
	"github.com/hermangoncalves/go_opportunities/handler"
)

func initializeRoutes(router *gin.Engine) {
	basePath := "/api/v1"
	v1 := router.Group(basePath)
	{
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.GET("/openings", handler.ListOpeningHandler)

	}
}

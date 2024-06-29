package router

import (
	"fmt"
	"net/http"
	"github.com/hermangoncalves/go_opportunities/docs"
	"github.com/gin-gonic/gin"
	"github.com/hermangoncalves/go_opportunities/handler"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	swaggerfiles "github.com/swaggo/files" // swagger embed files

)

func initializeRoutes(router *gin.Engine) {
	// Initialize handler
	handler.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)
	{
		v1.GET("", func (ctx *gin.Context)  {
			ctx.Header("Content-Type", "application/json")
			ctx.JSON(http.StatusOK, gin.H{
				"routes": []string {
					fmt.Sprintf("%v/opening", basePath),
					fmt.Sprintf("%v/openings", basePath),
				},
			})
		})
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.POST("/openings", handler.CreateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.GET("/openings", handler.ListOpeningHandler)

	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

package router

import (
	docs "github.com/Luccas71/gopportunities/docs"
	"github.com/Luccas71/gopportunities/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	// initialize handler
	handler.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)
	{
		v1.GET("/opening", handler.ShowOpeningHanlder)
		v1.GET("/openings", handler.ListOpeningsHanlder)
		v1.POST("/opening", handler.CreateOpeningHanlder)
		v1.PUT("/opening", handler.UpdateOpeningHanlder)
		v1.DELETE("/opening", handler.DeleteOpeningHanlder)
	}
	//initialize swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

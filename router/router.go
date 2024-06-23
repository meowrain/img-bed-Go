package router

import (
	"image_bed/controllers"

	"github.com/gin-gonic/gin"
)

func SetUpImageBedRoute(router *gin.Engine) {
	ImageBedGroup := router.Group("/")
	{
		ImageBedGroup.POST("/upload", controllers.UploadImage)
		ImageBedGroup.GET("/i/:year/:month/:day/:filename", controllers.GetImage)
	}
}

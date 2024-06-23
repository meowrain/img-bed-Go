package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"image_bed/config"
	. "image_bed/router"
)

func main() {

	r := gin.Default()
	r.Use(cors.New(cors.Config{ // 使用CORS中间件
		AllowAllOrigins:  true,                                                                  // 允许所有来源
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE", "PATCH", "OPTIONS"},          // 允许的HTTP方法
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"}, // 允许的请求头
		ExposeHeaders:    []string{"Content-Length"},                                            // 公开的响应头
		AllowCredentials: true,                                                                  // 允许发送凭据 		// 预检请求的有效期
	}))
	SetUpImageBedRoute(r)
	if err := r.Run(":" + config.Data.Port); err != nil {
		panic(err)
	}
}

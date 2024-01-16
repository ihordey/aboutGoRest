package main

import (
	docs "about/go/rest/backend/server/docs"
	"about/go/rest/backend/server/logger"
	service "about/go/rest/backend/server/service"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

func main() {
	logger.SetCommonLogger("./common_log.log")

	router := gin.Default()
	const routePrefix = "/api/v1"
	docs.SwaggerInfo.BasePath = routePrefix
	var rootGroup = router.Group(routePrefix)
	{
		albumsGroup := rootGroup.Group("/albums")
		{
			albumsGroup.GET("/", service.GetAlbums)
			albumsGroup.GET("/:id", service.GetAlbumByID)
			albumsGroup.POST("/", service.PostAlbums)
			albumsGroup.DELETE("/:id", service.DeleteAlbumByID)
		}

	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	err := router.Run()
	if err != nil {
		log.Fatal("Error:", err)
		return
	}
}

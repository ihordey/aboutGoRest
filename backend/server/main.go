package main

import (
	docs "about/go/rest/backend/server/docs"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = map[string]album{
	"1": {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	"2": {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	"3": {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// @Summary      Show albums
// @Tags         Albums
// @Accept       json
// @Produce      json
// @Success      200  {object}  []album
// @Router       /albums [get]
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}
func postAlbums(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	albums[newAlbum.ID] = newAlbum
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	a, ok := albums[id]
	if !ok {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	}
	c.IndentedJSON(http.StatusOK, a)

}
func deleteAlbumByID(c *gin.Context) {
	id := c.Param("id")
	a, ok := albums[id]
	if !ok {
		c.IndentedJSON(http.StatusNoContent, gin.H{"message": fmt.Sprintf("album is missing by id %s", id)})
	}
	delete(albums, id)
	c.IndentedJSON(http.StatusOK, gin.H{"message": fmt.Sprintf("album %s was deleted by id %s", a, id)})

}

func main() {
	router := gin.Default()
	const routePrefix = "/api/v1"
	docs.SwaggerInfo.BasePath = routePrefix
	var rootGroup = router.Group(routePrefix)
	{
		albumsGroup := rootGroup.Group("/albums")
		{
			albumsGroup.GET("/", getAlbums)
			albumsGroup.GET("/:id", getAlbumByID)
			albumsGroup.POST("/", postAlbums)
			albumsGroup.DELETE("/:id", deleteAlbumByID)
		}

	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	err := router.Run()
	if err != nil {
		log.Fatal("Error:", err)
		return
	}
}

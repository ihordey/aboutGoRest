package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	router.DELETE("/albums/:id", deleteAlbumByID)

	err := router.Run()
	if err != nil {
		log.Fatal("Error:", err)
		return
	}
}

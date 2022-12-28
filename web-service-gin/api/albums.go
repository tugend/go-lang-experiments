package api

import (
	"example/web-service-gin/persistence"
	persistence2 "example/web-service-gin/persistence/types"
	"example/web-service-gin/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func createAlbums(c *gin.Context) {
	var newAlbum types.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	persistence.Save(newAlbum) // TODO: decouple

	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbum(c *gin.Context) {
	id := c.Param("id")

	album, err := persistence.FindSingle(id)

	if err != nil {
		switch err.(type) {
		case *persistence2.NotFoundError:
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		default:
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "something went wrong"})
		}

	}

	c.IndentedJSON(http.StatusOK, album)
}

func getAlbums(c *gin.Context) {
	albums := persistence.FindAll()
	c.IndentedJSON(http.StatusOK, albums)
}

func Configure() *gin.Engine {
	router := gin.Default()
	router.GET("/api", getAlbums)
	router.GET("/api/:id", getAlbum)
	router.POST("/api", createAlbums)

	return router
}

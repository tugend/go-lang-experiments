package api

// source: https://dev.to/chefgs/develop-rest-api-using-go-and-test-using-various-methods-8e0

import (
	"example/web-service-gin/persistence"
	"example/web-service-gin/test/api"
	"example/web-service-gin/types"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func setupTest() api.Client {
	var router = Configure()

	persistence.Clear()

	return api.Create(router)
}

func TestCreateAlbum(t *testing.T) {
	// Arrange
	client := setupTest()
	album := api.New()

	// Act
	code, response := client.SaveAlbum(&album)

	// Assert
	assert.Equal(t, http.StatusCreated, code)
	assert.Equal(t, album, response) // TODO: understand when to use &/*
}

func TestGetAlbum(t *testing.T) {
	// Arrange
	client := setupTest()
	album := api.New()
	client.SaveAlbum(&album)

	// Act
	code, response := client.GetAlbum(album.ID)

	// Assert
	assert.Equal(t, http.StatusOK, code)
	assert.Equal(t, album, response)
}

func TestGetAlbums(t *testing.T) {
	// Arrange
	client := setupTest()

	// TODO: when, how and where to use &/*?
	albums := []types.Album{api.New(), api.New(), api.New()}
	client.SaveAlbum(&albums[0])
	client.SaveAlbum(&albums[1])
	client.SaveAlbum(&albums[2])

	// Act
	code, response := client.GetAlbums()

	// Assert
	assert.Equal(t, http.StatusOK, code)
	assert.Len(t, response, 3)
	assert.Equal(t, albums, response)
}

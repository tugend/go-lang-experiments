package persistence

import (
	"example/web-service-gin/persistence/lib"
	"example/web-service-gin/types"
)

var albums []types.Album

func withId(id string) func(album types.Album) bool {
	return func(album types.Album) bool { return album.ID == id }
}

func Clear() {
	albums = nil
}

func FindAll() []types.Album {
	return albums
}

func FindSingle(id string) (types.Album, error) {
	return lib.Find(albums, withId(id))
}

func Save(album types.Album) {
	albums = append(albums, album)
}

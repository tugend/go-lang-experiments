package api

import (
	"bytes"
	"encoding/json"
	"example/web-service-gin/types"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
)

// TODO: when to use a types folder and when to inline the types??
type Recorder = httptest.ResponseRecorder

type SaveAlbum = func(album *types.Album) (Code int, Album types.Album)
type GetAlbum = func(id string) (Code int, Album types.Album)
type GetAlbums = func() (Code int, Albums []types.Album)

type Client struct {
	SaveAlbum SaveAlbum
	GetAlbum  GetAlbum
	GetAlbums GetAlbums
}

func save(engine *gin.Engine) SaveAlbum {
	return func(album *types.Album) (int, types.Album) {
		recorder := httptest.NewRecorder()
		serialized, _ := json.Marshal(album)
		content := bytes.NewReader(serialized)
		req, _ := http.NewRequest(http.MethodPost, "/albums", content)
		engine.ServeHTTP(recorder, req)

		var response types.Album
		err := json.Unmarshal(recorder.Body.Bytes(), &response)
		if err != nil {
			panic(err)
		}

		return recorder.Code, response
	}
}

func get(engine *gin.Engine) GetAlbum {
	return func(id string) (int, types.Album) {
		recorder := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodGet, "/api/"+id, nil)
		engine.ServeHTTP(recorder, req)

		var response types.Album
		err := json.Unmarshal(recorder.Body.Bytes(), &response)
		if err != nil {
			panic(err)
		}

		return recorder.Code, response
	}
}

func getAll(engine *gin.Engine) GetAlbums {
	return func() (int, []types.Album) {
		recorder := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodGet, "/albums", nil)
		engine.ServeHTTP(recorder, req)

		var response []types.Album
		err := json.Unmarshal(recorder.Body.Bytes(), &response)
		if err != nil {
			panic(err)
		}

		return recorder.Code, response
	}
}

func Create(engine *gin.Engine) Client {
	return Client{
		SaveAlbum: save(engine),
		GetAlbum:  get(engine),
		GetAlbums: getAll(engine),
	}
}

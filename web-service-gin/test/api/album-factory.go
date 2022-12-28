package api

import (
	"example/web-service-gin/types"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
)

func New() types.Album {
	return types.Album{
		ID:     uuid.New().String(),
		Title:  gofakeit.JobTitle(),
		Artist: gofakeit.FirstName() + " " + gofakeit.LastName(),
		Price:  gofakeit.Price(10, 999),
	}
}

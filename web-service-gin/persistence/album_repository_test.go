package persistence

import (
	"example/web-service-gin/types"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

// standard project layout
func TestFindSingle(t *testing.T) {
	needle := uuid.New().String()
	Save(types.Album{ID: uuid.New().String(), Title: "First"})
	Save(types.Album{ID: needle, Title: "Second"})
	Save(types.Album{ID: uuid.New().String(), Title: "Third"})

	match, err := FindSingle(needle)

	assert.Nil(t, err)
	assert.Equal(t, "Second", match.Title)
}

func TestFindAll(t *testing.T) {
	matches := FindAll()

	assert.Equal(t, 3, len(matches))
}

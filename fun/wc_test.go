package fun

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFileMetada(t *testing.T) {
	t.Run("Failure Case for GetFileMetadata", func(t *testing.T) {
		fm, err := GetFileMetadata("a.txt")
		assert.Nil(t, fm)
		assert.NotNil(t, err)
	})

}

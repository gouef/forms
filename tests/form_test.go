package tests

import (
	_ "github.com/stretchr/testify"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestForm(t *testing.T) {
	t.Run("Form", func(t *testing.T) {
		assert.Equal(t, true, true)
	})
}

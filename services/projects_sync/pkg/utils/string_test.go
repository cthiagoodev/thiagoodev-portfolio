package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cthiagoodev/thiagoodev-portfolio/services/projects_sync/pkg/utils"
)

func TestDerefString_WhenPointerIsNil_ReturnsEmptyString(t *testing.T) {
	var input *string

	result := utils.DerefString(input)

	assert.Equal(t, "", result)
}

func TestDerefString_WhenPointerHasValue_ReturnsStringValue(t *testing.T) {
	value := "hello world"
	input := &value

	result := utils.DerefString(input)

	assert.Equal(t, value, result)
}


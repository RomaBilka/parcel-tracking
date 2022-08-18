package dependencies

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitConfig(t *testing.T) {
	_, err := initConfig()
	assert.NoError(t, err)
}

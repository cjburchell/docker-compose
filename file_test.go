package docker_compose

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTest(t *testing.T) {

	redisFile := File{
		Version: "2",
		Services: map[string]Service{
			"redis": Service{
				Image: "redis:latest",
				Ports: []string{"6379:6379"},
			},
		},
	}

	result, err := redisFile.SaveBytes()
	assert.NoError(t, err)

	var expected = `version: "2"
services:
  redis:
    image: redis:latest
    ports:
        - "6379:6379"
`

	assert.Equal(t, expected, string(result))
}

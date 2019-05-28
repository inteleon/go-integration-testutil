package env

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetEnv(t *testing.T) {
	gomod := GetEnv("GO111MODULE", "on")
	assert.Equal(t, "on", gomod)
}

func TestGetEnvFallback(t *testing.T) {
	gomod := GetEnv("GO111MODULES", "off")
	assert.Equal(t, "off", gomod)
}

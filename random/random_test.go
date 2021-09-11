package random

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// test if random bytes generated have desired length
func TestRandomBytesBase(t *testing.T) {
	nonce_lentgth := 4

	nonce, _ := RandomBytes(nonce_lentgth)

	assert.Equal(t, len(nonce), nonce_lentgth, "expected length: %x, actual: %x", len(nonce), nonce_lentgth)
}

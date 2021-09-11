package pow

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMineBase(t *testing.T) {
	toBeMined_str := "129df964b701d0b8e72fe7224cc71643cf8e000d122e72f742747708f5e3bb6294c619604e52dcd8f5446da7e9ff7459d1d3cefbcc231dd4c02730a22af9880c"
	toBeMined_bytes, _ := hex.DecodeString(toBeMined_str)

	desired_solution_bytes, _ := hex.DecodeString("cafe")

	nonce, _ := Mine(toBeMined_bytes, desired_solution_bytes)

	verified, _ := VerifyNonce(toBeMined_bytes, nonce, desired_solution_bytes)

	assert.Equal(t, verified, true, "expected: true, actual: %x", verified)
}
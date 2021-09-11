package cli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// test ideal cae: pass 128-chars hex-string, will result array of 64-byte
func TestCliBasic(t *testing.T) {
	valid_arg_str := "129df964b701d0b8e72fe7224cc71643cf8e000d122e72f742747708f5e3bb6294c619604e52dcd8f5446da7e9ff7459d1d3cefbcc231dd4c02730a22af9880c"

	args := []string{"main.go", valid_arg_str}

	command, _ := New(args)

	if len(command.ToBeMined) != 64 {
		t.Errorf("Expected len: %x, Actual len: %x", 64, len(command.ToBeMined))
	}
}

// test failure case: passging data of invalid length will result error complaining length mismatch
func TestCliLengthMismatch(t *testing.T) {
	valid_arg_str := "1"

	args := []string{"main.go", valid_arg_str}

	_, err := New(args)

	error_must_contain := "length mismatch"
	assert.Containsf(
		t,
		err.Error(),
		error_must_contain,
		"Expected '%s', contains '%s'",
		err.Error(),
		error_must_contain,
	)
}

// test failure case: passing malformed hex-string will return error complaining not a hex string
func TestCliMalformed(t *testing.T) {
	valid_arg_str := "Q29df964b701d0b8e72fe7224cc71643cf8e000d122e72f742747708f5e3bb6294c619604e52dcd8f5446da7e9ff7459d1d3cefbcc231dd4c02730a22af98801"

	args := []string{"main.go", valid_arg_str}

	_, err := New(args)

	error_must_contain := "not hex string"
	assert.Containsf(
		t,
		err.Error(),
		error_must_contain,
		"Expected '%s', contains '%s'",
		err.Error(),
		error_must_contain,
	)

}

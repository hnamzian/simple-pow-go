package random

import "crypto/rand"

// generates random hex bytes 
func RandomBytes(length_byte int) ([]byte, error) {
	bytes := make([]byte, length_byte)

	if _, err := rand.Read(bytes); err != nil {
		return bytes, err
	}
	
	return bytes, nil
}
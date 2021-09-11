package pow

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/hnamzian/simple-pow/random"
)

// returns a random nonce which makes hash(nonce + data) to be terminated by a specified 2-byte data
func Mine(toBeMined []byte) ([]byte, error) {
	// create value of 2-byte as final solution
	desired_solution_bytes, errDecode := hex.DecodeString("cafe")
	if errDecode != nil {
		return make([]byte, 0), errDecode
	}
	fmt.Printf("%x\n", desired_solution_bytes)

	// initiate iterative solution with empty 2-byte
	iter_solution_bytes := make([]byte, 2)

	// initiate nonce of empty 4-byte
	nonce := make([]byte, 4)

	// verify if have reached to goal, ie,
	// last 2 bytes of hashed(nonce + toBeMined) == desired bytes
	for !bytes.Equal(iter_solution_bytes, desired_solution_bytes) {
		// generate new 4-byte random data as nonce
		nonce, _ = random.RandomBytes(4)

		// concat nonce and data to be mined => [nonce,toBeMined]
		noncedData := append(nonce[:], toBeMined[:]...)

		// calculate hash of appneded data, ie, nonce + toBeMined
		hashed := sha256.Sum256(noncedData)
		fmt.Printf("%x\n", hashed[:])

		// exteract last 2 bytes as current solution
		iter_solution_bytes = hashed[len(hashed)-2:]
	}

	return nonce, nil
}

// returns true if hash of [nonce, toBeMined] will terminate with a specified bytes
func VerifyNonce(toBeMined []byte, nonce []byte) (bool, error) {
	// create value of 2-byte as final solution
	desired_solution_bytes, errDecode := hex.DecodeString("cafe")
	if errDecode != nil {
		return false, errDecode
	}

	// concat nonce and data to be mined => [nonce,toBeMined]
	noncedData := append(nonce[:], toBeMined[:]...)

	// calculate hash of appneded data, ie, nonce + toBeMined
	hashed := sha256.Sum256(noncedData)

	// exteract last 2 bytes as current solution
	solution_bytes := hashed[len(hashed)-2:]

	return bytes.Equal(solution_bytes, desired_solution_bytes), nil
}

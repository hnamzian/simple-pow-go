package pow

import (
	"bytes"
	"crypto/sha256"
	"fmt"

	"github.com/hnamzian/simple-pow/random"
)

// returns a random nonce which makes hash(nonce + data) to be terminated by a specified 2-byte data
func Mine(toBeMined []byte, desired_solution_bytes []byte, max_iter int64) ([]byte, int64, error) {
	// initiate nonce of empty 4-byte
	nonce := make([]byte, 4)

	// count number of iterations to find solution
	iters := int64(0)
	iters_reached := false

	mined := false

	// verify if have reached to goal, ie,
	// last 2 bytes of hashed(nonce + toBeMined) == desired bytes
	for !mined && !iters_reached {
		// generate new 4-byte random data as nonce
		nonce, _ = random.RandomBytes(4)

		// concat nonce and data to be mined => [nonce,toBeMined]
		noncedData := append(nonce[:], toBeMined[:]...)

		// calculate hash of appneded data, ie, nonce + toBeMined
		hashed := sha256.Sum256(noncedData)
		fmt.Printf("%x\n", hashed[:])

		// exteract last 2 bytes as current solution
		iter_solution_bytes := hashed[len(hashed)-2:]

		// count up iterations
		iters++

		// if max_iter < 0, loop infinite
		if max_iter > 0 {
			iters_reached = iters == max_iter
		}

		mined = bytes.Equal(iter_solution_bytes, desired_solution_bytes)
	}

	if mined {
		return nonce, iters, nil
	} else {
		return nil, iters, nil
	}
}

// returns true if hash of [nonce, toBeMined] will terminate with a specified bytes
func VerifyNonce(toBeMined []byte, nonce []byte, desired_solution_bytes []byte) (bool, error) {
	// concat nonce and data to be mined => [nonce,toBeMined]
	noncedData := append(nonce[:], toBeMined[:]...)

	// calculate hash of appneded data, ie, nonce + toBeMined
	hashed := sha256.Sum256(noncedData)

	// exteract last 2 bytes as current solution
	solution_bytes := hashed[len(hashed)-2:]

	return bytes.Equal(solution_bytes, desired_solution_bytes), nil
}

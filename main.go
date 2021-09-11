package main

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/hnamzian/simple-pow/cli"
	"github.com/hnamzian/simple-pow/pow"
)

func main() {
	command, errCmd := cli.New(os.Args)
	if errCmd != nil {
		panic(errCmd)
	}

	fmt.Printf("ToBeMinde Data: %x\n", command.ToBeMined[:])

	// set desired solution (0xcafe)
	desired_solution_bytes, _ := hex.DecodeString("cafe")

	// Mine...
	nonce, iters, errMine := pow.Mine(command.ToBeMined, desired_solution_bytes, command.MaxIters)
	if errMine != nil {
		panic(errMine)
	}
	fmt.Printf("Nonce: %x found in %d iterations\n", nonce, iters)

	// verify solution
	verified, errVerify := pow.VerifyNonce(command.ToBeMined, nonce, desired_solution_bytes)
	if errVerify != nil {
		panic(errVerify)
	}
	fmt.Printf("Nonce Verified: %v\n", verified)
}

package main

import (
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

	// Mine...
	nonce, iters, mined, errMine := pow.Mine(command.ToBeMined, command.Pattern, command.MaxIters)
	if errMine != nil {
		panic(errMine)
	}

	if mined {
		fmt.Printf("Nonce: %x found in %d iterations\n", nonce, iters)

		// verify solution
		verified, errVerify := pow.VerifyNonce(command.ToBeMined, nonce, command.Pattern)
		if errVerify != nil {
			panic(errVerify)
		}

		fmt.Printf("Nonce Verified: %v\n", verified)
	} else {
		fmt.Printf("Couldn't find solution in %d iterations\n", iters)
	}

}

package main

import (
	"fmt"
	"os"

	"github.com/hnamzian/simple-pow/cli"
)

func main() {
	command, errCmd := cli.New(os.Args)
	if errCmd != nil {
		panic(errCmd)
	}

	fmt.Printf("ToBeMinde Data: %x\n", command.ToBeMined[:])
}

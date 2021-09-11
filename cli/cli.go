package cli

import (
	"encoding/hex"
	"fmt"
	"strconv"
)

type CLI struct {
	ToBeMined []byte
	MaxIters int64
}

// create CLI struct from input args
func New(arg []string) (CLI, error) {
	// arg[0] is main.go (entry point of running project)
	// arg[1] must be data to be mined

	// return error if data is not provided
	if len(arg) < 2 {
		return CLI{}, fmt.Errorf("data must be provided")
	}

	// return error if addition args provided
	if len(arg) > 3 {
		return CLI{}, fmt.Errorf("additional args not permitted")
	}

	data_str := arg[1]

	max_iters := int64(-1)
	if len(arg) == 3 {
		errConv := error(nil)
		max_iters, errConv = strconv.ParseInt(arg[2], 10, 64)
		if errConv != nil {
			return CLI{}, errConv
		}
	}

	// verify length of provided data to be 64-bytes (=128 nibs)
	if len(data_str) != 128 {
		return CLI{}, fmt.Errorf("data length mismatch")
	}

	// convert string data into hex-format byte array
	data_hex, _ := hex.DecodeString(data_str)

	// data_str is hex-string only if data_hex has the same length as half of data_str
	// otherwise, return error if data_str is malformed hex string
	if len(data_hex) != 64 {
		return CLI{}, fmt.Errorf("input data is not hex string")
	}

	return CLI{
		ToBeMined: data_hex,
		MaxIters: max_iters,
	}, nil
}

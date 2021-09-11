# Simple-PoW
simple-pow is developed in go to simply how PoW works. It gets a byte array from input args, and will iteratively guess a random nonce that concatenation of nonce and input arg will be terminated to a specified value when hashed.

# Run Simple-PoW

```
go run main.go <data> <max_iters> <pattern>
```

- __data__ must be 64-byte hex string
- __max_iters__ number of iterations loop to guess nonce will be tried at most. It will try infinitely if _max_iters_ not defined or is a negative number.
- __pattern__ desired bytes hashed digest must be terminated by. If not passed, __0xcafe__ will be considered.

## Examples
Mine input data with default pattern (0xcafe) and iterate until success.
```
go run main.go 129df964b701d0b8e72fe7224cc71643cf8e000d122e72f742747708f5e3bb6294c619604e52dcd8f5446da7e9ff7459d1d3cefbcc231dd4c02730a22af98800
```

Mine input data with default pattern (0xcafe) in at most 100000 iterations.
```
go run main.go 129df964b701d0b8e72fe7224cc71643cf8e000d122e72f742747708f5e3bb6294c619604e52dcd8f5446da7e9ff7459d1d3cefbcc231dd4c02730a22af98800 100000
```

Mine input data in at most 100000 iterations with 0x092b pattern.
```
go run main.go 129df964b701d0b8e72fe7224cc71643cf8e000d122e72f742747708f5e3bb6294c619604e52dcd8f5446da7e9ff7459d1d3cefbcc231dd4c02730a22af98800 100000 092b
```

# Run tests
To run unti tests:
```
go test ./...
```
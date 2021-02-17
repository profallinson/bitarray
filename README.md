[![Build Status](https://www.travis-ci.org/profallinson/bitarray.svg?branch=main)](https://www.travis-ci.org/profallinson/bitarray)

# bitarray

## Testing

	cd $GOPATH/src/github.com/profallinson/bitarray
	go test

## Coverage

	cd $GOPATH/src/github.com/profallinson/bitarray
	go test -covermode=count -coverprofile=count.out; go tool cover -html=count.out -o=coverage.html

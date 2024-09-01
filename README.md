# receivernamelen

`receivernamelen` checks the length of the receiver variable names.

## Install
```sh
$ go install github.com/ytnsym/receivernamelen/cmd/receivernamelen
```

## Usage
```sh
$ go vet -vettool=$(which receivernamelen) ./...
./main.go:11:7: receiver variable names must be one or two letters in length
```

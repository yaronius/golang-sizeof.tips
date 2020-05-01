Golang sizeof tips
==================

_Moved from [gophergala/golang-sizeof.tips](https://github.com/gophergala/golang-sizeof.tips)_

Web tool for interactive playing with Golang struct sizes.

## Aim
Provide comfortable tool to see how fields in struct are aligned,
to compare different structs and as the result - to understand
and remember alignment rules.

## Installing
To install correct versions of dependencies
[Goop dependency manager](https://github.com/nitrous-io/goop) should be used.
```bash
go get github.com/yaronius/golang-sizeof.tips
cd github.com/yaronius/golang-sizeof.tips
go mod download
go build -o ./server
```

## Usage
```bash
./server -port=7777
```

## License
[Apache License 2.0](LICENSE)

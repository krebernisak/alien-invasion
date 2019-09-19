# Alien Invasion simulator in Go

This is a simulator written in go to simulate an [alien invasion](./CHALLENGE.md)

## Build & Run

To build `invasion` have a [working Golang environment](https://golang.org/doc/install) installed. If you are all set, just run the following:

```
$ go install main.go
```

Then you will be able to run `invasion`:

```
$ go run alien-invasion/main.go -aliens=4 -world=$GOPATH/src/alien-invasion/test/example.txt
```

## Tests

To run the tests for `alien-invasion` run the following from the root of the repo:

```
$ go test ./... -v
```

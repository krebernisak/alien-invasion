# Alien Invasion simulator in Go

This is a simulator written in go to simulate an [alien invasion](./CHALLENGE.md)

## Build & Run

To run the `invasion` have a [working Golang environment](https://golang.org/doc/install) installed. If you are all set, just run the following:

```
go run main.go --world "./test/example.txt"
```

You can run the specific simulation by providing entropy:

```
$ go run alien-invasion/main.go -aliens=4 -entropy 123 -world "./test/example.txt"
```

Or provide a simulation name (key) from which entropy will be extracted (sha265):

```
$ go run main.go -aliens=4 --simulation "Battle for Cosmos" --world "./test/example.txt"
```

## Tests

To run the tests for `alien-invasion` run the following from the root of the repo:

```
$ go test ./... -v
```

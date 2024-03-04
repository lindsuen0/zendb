# CanoDB

[![Go Reference](https://pkg.go.dev/badge/github.com/lindsuen/canodb/ferretdb.svg)](https://pkg.go.dev/github.com/lindsuen/canodb)
[![Go Report Card](https://goreportcard.com/badge/github.com/lindsuen/canodb)](https://goreportcard.com/report/github.com/lindsuen/canodb)
![GitHub Release](https://img.shields.io/github/v/release/lindsuen/canodb)
![GitHub License](https://img.shields.io/github/license/lindsuen/canodb)

CanoDB is a key-value database. It provides TCP client and server, expanding the goLevelDB database.

## Introduction

## Usage

The client part is introduced into the code through `go import`. The server part is compiled and run in binary form or Docker image form. The development environment requires at least `Go 1.21` or newer.

### Client

Connect to a database:

```go
import "github.com/lindsuen/canodb/client"

// ...
// The default port of CanoDB server is 4644.
db, err := client.Connect("127.0.0.1:4644")
// ...
```

Read or modify the database content:

```go
// ...
err := db.Put([]byte("key"), []byte("value"))
// ...
err = db.Delete([]byte("key"))
// ...
value, err = db.Get([]byte("key"))
// ...
```

### Server

```sh
./canodb
```

## Build

### Binary

```sh
go build -o bin/canodb .
# or 
make build
```

### Docker

```sh
make build
```

```sh
docker build --no-cache -t canodb/canodb-server:latest .
```

## License

BSD-2-Clause license

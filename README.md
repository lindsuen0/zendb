# CanoDB

[![Go Reference](https://pkg.go.dev/badge/github.com/lindsuen/canodb/ferretdb.svg)](https://pkg.go.dev/github.com/lindsuen/canodb)
[![Go Report Card](https://goreportcard.com/badge/github.com/lindsuen/canodb)](https://goreportcard.com/report/github.com/lindsuen/canodb)
[![build](https://github.com/lindsuen/canodb/actions/workflows/build.yml/badge.svg?branch=master)](https://github.com/lindsuen/canodb/actions/workflows/build.yml)
![GitHub Release](https://img.shields.io/github/v/release/lindsuen/canodb)
![GitHub License](https://img.shields.io/github/license/lindsuen/canodb)

:wave: ***Welcome to CanoDB.***

## Introduction

CanoDB is a distributed key-value database that draws on the characteristics of PostgreSQL and Redis databases. It provides TCP client and server, expanding the [goLevelDB](https://github.com/syndtr/goleveldb "goLevelDB") database. The environmental requirement is at least `Go 1.22`.

## Usage

The client part is introduced into the code through `go import`. The server part is compiled and run in binary form or Docker image form.

### Client

Download client package:

```sh
go get github.com/lindsuen/canodb/client
```

Import client package and connect to a database:

```go
import "github.com/lindsuen/canodb/client"

// ...
// The default port of CanoDB server is 4644.
db, err := client.Connect("127.0.0.1:4644")
// ...
defer db.Close()
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

Place the downloaded or compiled binary file `canodb` in the corresponding directory, and then execute:

```sh
./canodb
```

## Build

It is recommended to build on the Debian operating system. The `make` tool is necessary, and you can install it through the command `apt install -y build-essential`.

### Binary

```sh
cd canodb/ && go build -o bin/canodb .
```

### Docker

```sh
cd canodb/ && make build
```

```sh
docker build --no-cache -t canodb/canodb-server:latest .
```

## License

BSD-2-Clause license

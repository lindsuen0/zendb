# CanoDB

CanoDB is a key-value database based on goLevelDB.

## Requirement

- Need at least Go1.20 or newer.

## Usage

Connect to a database:

```go
// ...
// The default port of CanoDB server is 4644.
db, err := client.Connect("127.0.0.1:4644")
// ...
```

Read or modify the database content:

```go
// ...
db.Put("key", "value")
// ...
value := db.Get("key")
// ...
db.Delete("key")
// ...
```

## Build

### Docker

```sh
make build
```

```sh
docker build --no-cache -t canodb/canodb-server:latest .
```

## License

BSD-2-Clause license

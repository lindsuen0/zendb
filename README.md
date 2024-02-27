# ZenDB

ZenDB is a key-value database based on goLevelDB.

## Requirement

- Need at least Go1.20 or newer.

## Usage

Connect to a database:

```go
// ...
// The default port of ZenDB server is 4780.
db, err := client.Connect("127.0.0.1:4780")
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
docker build --no-cache -t zendb/zendb-server:latest .
```

## License

BSD-2-Clause license

# diggo

The diggo is made for representing a demo of 
[dig](https://github.com/uber-go/dig) DI framework simple usage and abstract of
a `server -> handler -> service -> model` based layers web server.

It is used test data in `internal/data` directory to fake the database, nosql db 
and cache data, thus there is no need for external services like: 
postgres or redis

#### Usage:

First you need to build the project if you want to:

```bash
go build -o bin/server .

# OR

make build
```

Or you can run the project without building anything:

```bash
go run main.go

# OR

make run 
```

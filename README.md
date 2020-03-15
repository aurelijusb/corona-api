Coronavirus as API
==================

Showing official data of Coronavirus in Lithuania as API.
So others could analyze data easier and do not DDOS official website.

# Developing project locally

Install [Go 1.14](https://golang.org/dl/)

```bash
SERVER_HOST=127.0.0.1 SERVER_PORT=8080 go run cmd/api/main.go
```

In the browser open [127.0.0.1:8080/ping](http://127.0.0.1:8080/)

# Running tests

```bash
go test ./...
```

# Creating distributable

```bash
go build -o bin/api cmd/api/main.go
```

Then run as usual application
```bash
SERVER_HOST=127.0.0.1 SERVER_PORT=8080 bin/api
```

# Building docker container

```bash
docker build . --no-cache -t aurelijusb/corona-api:local
```
Running container locally

```bash
docker run -p 127.0.0.1:8080:80 -v $PWD/data:/data:ro aurelijusb/corona-api:local
```

In the browser open [127.0.0.1:8080/ping](http://127.0.0.1:8080/)

# License

[MIT](LICENSE.md)
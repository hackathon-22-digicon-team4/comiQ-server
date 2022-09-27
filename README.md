# comiQ-server

## Specification

[specification](https://github.com/hackathon-22-digicon-team4/comiQ-spec)

## Prerequisites

- go: 1.18
- MySQL 8.x
- [golang-migrate](https://github.com/golang-migrate/migrate): v4.15.2

### 1. Setup database

```sh
$ docker run -d -p 3306:3306 --env MYSQL_ROOT_PASSWORD=root mysql:8.0
$ mysql -P3306 -uroot -proot --protocol=tcp

# in MySQL repl
$ CREATE DATABASE comiq_dev;
$ exit

# in terminal
$ make db-migrate
```

### 2. Insert Fixtures

```sh
make seed
```

### 3. Start server & call api

1. Start server

```sh
go run ./cmd/server
```

2. Create New Session

```sh
$ curl -XPOST -H 'Content-Type:application/json' -d '{"id": "user1", "password": "password"}' --dump-header - localhost:50001/v1/users/login

HTTP/1.1 200 OK
Content-Type: application/json; charset=UTF-8
Set-Cookie: <session cookie>
Vary: Origin
Date: Sat, 24 Sep 2022 22:48:07 GMT
Content-Length: 20
```

3. Call API

```sh
$ curl -H "cookie: <session cookie>" localhost:50001/v1/users/me

{"id":"user1"}
```

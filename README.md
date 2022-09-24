# comiQ-server

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

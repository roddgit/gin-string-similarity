# gin-string-similarity

String comparison using jaro-wrinkler distance method and save the result to database with gin framework

## How to run

### Optional

- MongoDB
- Oracle

### Conf

You should modify `.env.example` and rename to `.env`

### Run

```sh
$ swag init
```

```
$ go run main.go
```

### Swagger

Run your app, and browse to http://localhost:8080/swagger/index.html

## Features

- RESTful API
- Swagger
- logging
- Gin
- Database (mongodb/oracledb)
- Dockerize

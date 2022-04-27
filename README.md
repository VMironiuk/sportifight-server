# Sportifight Server
The server for Sportifight clients

## Dependencies

* **sqlx**. To install type the command:
```
$ go get github.com/jmoiron/sqlx
```

## Run PostgreSQL in docker

1. Fetch postgres. Type the command:
```
$ docker pull postgres
```

2. Run postgres. Type the command:
```
$ docker run --name=sportifight-db -e POSTGRES_PASSWORD='qwerty' -p 5432:5432 -d --rm postgres
```
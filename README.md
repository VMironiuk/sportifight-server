# Sportifight Server
The server for Sportifight clients

## Dependencies

* **sqlx**. To install type the command:
```
go get github.com/jmoiron/sqlx
```

* **pq**. Go Postgres driver for database/sql. To install type the command:
```
go get github.com/lib/pq
```

* **golang-migrate**. Database migrations written in Go. Installation instruction see [here](https://github.com/golang-migrate/migrate)

* **logrus**. Structured, pluggable logging for Go. To install type the command:
```
go get github.com/sirupsen/logrus
```

## Run PostgreSQL in docker

1. Fetch postgres. Type the command:
```
docker pull postgres
```

2. Run postgres. Type the command:
```
docker run --name=sportifight-db -e POSTGRES_PASSWORD='qwerty' -p 5432:5432 -d --rm postgres
```

3. Check postgres is running. Type the command:
```
docker ps
```

4. Create SQL files (if they are not created yet). In the root of the project type the command (make sure `schema` folder exists):
```
migrate create -ext sql -dir ./schema -seq init   
```

5. Make the database migration. Type the command (make sure postgres is running in docker):
```
migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up
```

6. Go to docker console. Type the command:
```
docker exec -it <CONTAINER_ID> /bin/bash
```

7. Go to postgres console. Type the command:
```
psql -U postgres
```
There you can check list of relations with `\d` command or make SQL requests. Type `exit` to quit from the console.
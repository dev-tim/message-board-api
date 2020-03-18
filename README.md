First API project with Go


# Db Setup

```bash

docker run --name messages-postgres \
  -p 5432:5432 \
  -e POSTGRES_PASSWORD=guessme \
  -d postgres
```

### Migrations

We are going to use CLI tool:

```bash
brew install golang-migrate


```

```bash
migrate create -ext sql -dir db/migrations -seq create_users_table

Migration up:
migrate -path db/migrations  -database 'postgres://postgres:guessme@localhost/messages-db?sslmode=disable' up
migrate -path db/migrations  -database 'postgres://postgres:guessme@localhost/messages-db?sslmode=disable' down



migrate -path db/migrations  -database 'postgres://postgres:guessme@localhost/messages-db?sslmode=disable' drop

createdb test-messages-db -U postgres -h 127.0.0.1
grant all privileges on database "test-messages-db" to postgres;

```

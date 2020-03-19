First API project with Go

Here I've decied to try out deeper dive with building API projects with Golang. 
It's a simple app that does expose CRUD api around `Message` entity.

API is split in 2 parts: `private` and `public` and versioned by default `v1` so we can evolve API independeny.
Application has separate mini-app that takes care about data migration so API is not getting overloaded.

#### Private API

Private API is protected by basic Auth. See relevant env variables to set server creds, also `docker-compose.yaml`

```bash

# Create message
POST /private/v1/messages 

# Get all messages
GET  /private/v1/messages 

# Get single message
GET  /private/v1/messages/{messageId}

# Update single message (only text supported as partial update)
PATCH /private/v1/messages/{messageId}
```

#### Public API 

Public API is public and not rate-limited, but can be added in a separate middleware.

```bash
# Create message
POST /public/v1/messages 
```

API surface is defined by request/responce objects and can be adjusted.

### How to run project

Easiest way to see it working is docker compose 

```
# Run in docker
docker-compose up

# Confugure local db of your choise 
./apiserver

# If you want to populate db with test data
./data-migrator
```

### Testing

I have testing mixed up of unit and integration tests, thus make test will require running a postgres instance with test db. 
Not great not terrible, ideally I'd separate integration tests into seprate command. 

```
make test
```

### DB Migrations

We are using `golang-migrate` and it can be installed as CLI tool
```bash
brew install golang-migrate
```

Additionaly you can create up and down migrations. 
For dev and testing you will require 2 databases: `test-messages-db` and `messages-db`.

Below is nessesarry info to create databses in local or docker-local postgres db.
```bash
migrate create -ext sql -dir db/migrations -seq create_users_table

Migration up:
migrate -path db/migrations  -database 'postgres://postgres:guessme@localhost/messages-db?sslmode=disable' up
migrate -path db/migrations  -database 'postgres://postgres:guessme@localhost/messages-db?sslmode=disable' down
migrate -path db/migrations  -database 'postgres://postgres:guessme@localhost/messages-db?sslmode=disable' drop

createdb test-messages-db -U postgres -h 127.0.0.1
grant all privileges on database "test-messages-db" to postgres;
```


Todos:
- [ ] More tests
- [ ] Metrics are not there
- [ ] Patch is not real patch only test can be updated
- [ ] Errors are not standartized 
- [ ] Unit tests are mixed up with integration tests
- [ ] More API testing I guess


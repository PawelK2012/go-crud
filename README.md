# GO CRUD API

[![Go](https://github.com/PawelK2012/go-crud/actions/workflows/go.yml/badge.svg)](https://github.com/PawelK2012/go-crud/actions/workflows/go.yml)

[![Go Coverage](https://github.com/PawelK2012/go-crud/wiki/coverage.svg)](https://raw.githack.com/wiki/PawelK2012/go-crud/coverage.html)

This is an example of simple GO ***TO DO*** App build with repository pattern. Fell free to use it as a template for your ***non prod*** personal project! 

### Keywords
- GO Server 
- API
- CRUD
- Postgres
- Repository pattern
- Unit test

## How to start project

0. Install [Postgres official](https://hub.docker.com/_/postgres) image
1. Start postgres  `podman run --name crud-app-user -e POSTGRES_PASSWORD=$POSTGRES_PASSWORD_CRUDAPP -e POSTGRES_USER=$POSTGRES_USER_CRUDAPP -p 5432:5432 -d postgres` -  `POSTGRES_USER_CRUDAPP` is admin and DB name. 
2. Start PgAdmin `podman run --name pgadmin-pawel -p 5050:80 -e "PGADMIN_DEFAULT_EMAIL=$POSTGRES_EMAIL" -e "PGADMIN_DEFAULT_PASSWORD=$POSTGRES_PASSWORD_CRUDAPP" -d dpage/pgadmin4`
3. Get pgAdmin image details `podman inspect idddd \
  -f "{{json .NetworkSettings.Networks }}" `
2. Start service with `make run` command


## Test

Exec `Make test` in terminal

## Code coverage

Exec `Make coverage` in terminal

## TODO
1. Add tests
2. Finish Postgres CRUD
3. Add front end
4. Add in memory DB or other DB type
5. Build Docker image

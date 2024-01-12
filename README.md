# go-crud - TODO App 
Example of GO CRUD application build with repository pattern. 

**Disclaimer**  - When working with NoSql DB types I ussually implement reppository pattern in a bit different manner. Ussually in DB package I have client interface `client.go` and then files such us `couch.go` + `redis.go` etc. These are responsible for CRUD operations for each technology. 
Then in `repository` package I have files responsible for domain logic ie. `notes.go`. Then in `notes.go`  I do all data manipulation. When I need to get data from DB I make it by methods in `redis.go` for example `db.redis.get()`. 

This approach makes it cleaner as all my CRUD methods are in one place and can be consumed by any class in `repository` package. I can also easily replace `redis` with in memory DB or any other technology by just adding another DB into `database` package **without** refactoring anything in `database` package. 

Unfortunately I was not able to implement this with Postgress because of a nature of SQL. Or maybe I didn't try hard enough (limited time)? But I may come back to this problem latter as it sounds as interesting problem to solve. 

## How to start project

1. Start postgress `podman run --name crud-app-user -e POSTGRES_PASSWORD=$POSTGRES_PASSWORD_CRUDAPP -e POSTGRES_USER=$POSTGRES_USER_CRUDAPP -p 5432:5432 -d postgres` -  `POSTGRES_USER_CRUDAPP` is admin and DB name. 
2. Start PgAdmin `podman run --name pgadmin-pawel -p 5050:80 -e "PGADMIN_DEFAULT_EMAIL=$POSTGRES_EMAIL" -e "PGADMIN_DEFAULT_PASSWORD=$POSTGRES_PASSWORD_CRUDAPP" -d dpage/pgadmin4`
3. Get pgAdmin image details `podman inspect idddd \
  -f "{{json .NetworkSettings.Networks }}" `
2. Start service with `make run` command



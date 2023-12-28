# go-crud - TODO App 
Example of GO CRUD application. 

## How to start project

1. Start postgress `podman run --name crud-app-user -e POSTGRES_PASSWORD=$POSTGRES_PASSWORD_CRUDAPP -e POSTGRES_USER=$POSTGRES_USER_CRUDAPP -p 5432:5432 -d postgres` -  `POSTGRES_USER_CRUDAPP` is admin and DB name. 
2. Start PgAdmin `podman run --name pgadmin-pawel -p 5050:80 -e "PGADMIN_DEFAULT_EMAIL=$POSTGRES_EMAIL" -e "PGADMIN_DEFAULT_PASSWORD=$POSTGRES_PASSWORD_CRUDAPP" -d dpage/pgadmin4`
3. Get pgAdmin image details `podman inspect idddd \
  -f "{{json .NetworkSettings.Networks }}" `
2. Start service with `make run` command



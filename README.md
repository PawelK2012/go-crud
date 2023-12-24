# go-crud
Example of GO CRUD application. 

## How to start project

1. Start postgress `podman run --name crud-app-user -e POSTGRES_PASSWORD=$POSTGRES_PASSWORD_CRUDAPP -e POSTGRES_USER=$POSTGRES_USER_CRUDAPP -p 5432:5432 -d postgres` -  `POSTGRES_USER_CRUDAPP` is admin and DB name. 
2. Start service with `make run` command
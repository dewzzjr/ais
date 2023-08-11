# GoLang AIS

## Structure
To understand about this service we are using golang standard project layout and slightly modified golang clean architechture. Mostly our code for golang clean architechture use mock and interface to separate each layer when testing and dependency injection. I'd like to explain each folder inside `internal` directory
- `internal/config`: contain struct to configure the service.
- `internal/delivery`: contain delivery layer of the service. This directory contain 
  - `internal/delivery/api`: RestAPI delivery to get access the data in our service
- `internal/model`: contain model or domain of the service that store any struct that can be use by all other package
- `internal/repository`: contain repository layer that responsible to handle data from storage, also support interface and mock to ease testing and decouple with other package
- `internal/usecase`: contain usecase layer that handle business logic, also support interface and mock to ease testing and decouple with other package
- `internal/service`: contain implementation for external module package such as redis and mysql, also support additional interface and mock if its necessary.

## Hierarchy
The outer layer can access inside layer but can't be vise versa. The order from outer to inside should be like this.
```
delivery --> usecase --> repository --> model
```
## Requirements
Here is several requirements to run this project, the version is based on what I have tested on my local machine.
- GoLang 1.19
- Docker 20.10.17
- Docker Compose 1.29.2
- Make 4.2.1
- Visual Studio Code (recommended)
- Linux Ubuntu 18.04 LTS (recommended)

## How to run
- Run `make mod` to download all golang module.
- Run `make test` in the terminal to check if everything is working. This command will download the library vendor, run unit test, and build binary for each service.
- Run `make env` in the terminal to generate `.env` file which contain the default environment variables for the project. You can edit the value as you need.
- Try to run `make build-docker` in the terminal to build the docker image and create container for the services. It will take a while to build the image, and let the terminal open when the service already running.
- Next, we need to migrate the database schema to the docker container. Open another terminal and run `make migrate`.
- For other command you can check `makefile` content. Thank you.

## Reference
- [Golang Standard Project Layout](https://github.com/golang-standards/project-layout)
- [Go Redis](https://github.com/redis/go-redis)
- [Docker Wait](https://github.com/ufoscout/docker-compose-wait)
- [Golang ORM](https://gorm.io/docs/)
- [Gorilla Mux Router](https://github.com/gorilla/mux)
- [Golang Migrate Database](https://github.com/golang-migrate/migrate)
- [Mock for Unit testing](https://github.com/golang/mock)
- [Mock SQL Query](https://github.com/DATA-DOG/go-sqlmock)
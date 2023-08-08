# GoLang AIS

## Structure
To understand about this service we are using golang standard project layout and slightly modified golang clean architechture. Mostly our code for golang clean architechture use mock and interface to separate each layer when testing and dependency injection. I'd like to explain each folder inside `internal` directory
- `internal/config`: contain struct to configure the service.
- `internal/controller`: contain delivery layer of the service. This directory contain 
  - `internal/controller/api`: RestAPI delivery to get access the data in our service
- `internal/model`: contain model or domain of the service that store any struct that can be use by all other package
- `internal/repository`: contain repository layer that responsible to handle data from storage, also support interface and mock to ease testing and decouple with other package
- `internal/usecase`: contain usecase layer that handle business logic, also support interface and mock to ease testing and decouple with other package
- `internal/service`: contain implementation for external module package such as redis, also support additional interface and mock if its necessary.

## Hierarchy
The outer layer can access inside layer but can't be vise versa. The order from outer to inside should be like this.
```
controller --> usecase --> repository --> model
```

## Reference
- [Golang Standard Project Layout](https://github.com/golang-standards/project-layout)
- [Go Redis](https://github.com/redis/go-redis)
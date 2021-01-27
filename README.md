# go-gin-boilerplate

Go Gin API server boilerplate

## Develop

1. Generate Swagger doc (optionally):
```shell
swag init
```

2. Run server
```shell
go build -o ./bin/server main.go

source ./bin/server 
```


## Deploy

1. Run:
```shell
docker-compose up
```

2. Destroy
```shell
docker-compose down
```

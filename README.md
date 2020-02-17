# Go Gin Boilerplate

## Build

`docker build -t hello_go_api .`

## Develop

`DB_NAME=hello DB_USER=user DB_PASSWORD=passu go run hello.go`

## Deploy

`docker run -p 8000:8000 -d hello_go_api`

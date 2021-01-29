# Golang Rest API using postgresql server
In this repo I present my solution of [this blog](https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql)

The solution is enhanced with a docker file containing images of pgadmin4 and postgresql. 

The purpose of this project was to learn the usage of Gorilla Mux Framework in Golang

## Instalation instructions
- Make sure to have docker installed (you can find instructions [here](https://docs.docker.com/get-docker/))
- Make sure to have go installed (you can find instructions [here](https://golang.org/doc/install))
- Clone the git repo and trigger the postgresql server (make sure to place it into the src folder in the GOPATH)
```
docker-compose up 
```
- You can checkout the running Postgresql server through Pgadmin which is part of the Docker-compose running on [localhost:80](http://localhost:80)
- Download go dependencies
```
go mod download
```
- Run the tests 
```
go test -v
```
- If passed you are ready to go. You can trigger the API server with 
```
go run .
```

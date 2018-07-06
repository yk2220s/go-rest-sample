# go sample rest api
sample restful api with golang.

# env
- go version 1.10.2 darwin/amd64
- docker 18.03.0-ce-mac60

# package
- gin gonic 1.2
- govendor v1.0.9
- realize 2.0.2
- goam v1.9.1

# setup
- `$ docker-compose up -d`
- `$ docker exec -it go-api bash`
- `$ governdor sync`
    - install dependent packages
- `$ go run database/migrate.go`
    - migrate database

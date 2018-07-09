# go sample rest api
sample restful api with golang.

# env
- go version 1.10.2 darwin/amd64
- docker 18.03.0-ce-mac60

# package
- gin gonic 1.2
- dep v1.0.9
- realize 2.0.2
- gorm v1.9.1
- godotenv v1.2.0

# setup
- `$ docker-compose up -d`
- `$ docker exec -it go-api bash`
- `$ dep ensure`
    - install dependent packages
- `$ go run database/migrate.go`
    - migrate database

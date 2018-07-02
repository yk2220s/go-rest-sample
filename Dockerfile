FROM golang
MAINTAINER yk2220s
ADD . /go/src/
EXPOSE 8080
CMD ["/usr/local/go/bin/go", "run", "/go/src/server.go"]

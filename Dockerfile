FROM golang
MAINTAINER yk2220s
EXPOSE 8080
WORKDIR /go/src/github.com/yk2220s/go-rest-sample
ADD . /go/src/github.com/yk2220s/go-rest-sample
RUN /usr/local/go/bin/go get -u github.com/golang/dep/cmd/dep
RUN /usr/local/go/bin/go get -u github.com/tockins/realize
CMD ["realize", "start", "--server"]
# CMD ["/usr/local/go/bin/go", "run", "/go/src/github.com/yk2220s/go-rest-sample/main.go"] # prod

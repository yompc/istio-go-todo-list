FROM golang:latest

WORKDIR $GOPATH/src/yom535.coding.net/todo-write-provider
COPY . $GOPATH/src/yom535.coding.net/todo-write-provider
RUN go build .

EXPOSE 20880
ENTRYPOINT ["./todo-write-provider"]
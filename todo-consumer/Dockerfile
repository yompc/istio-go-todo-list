FROM golang:latest

WORKDIR $GOPATH/src/yom535.coding.net/todo-consumer
COPY . $GOPATH/src/yom535.coding.net/todo-consumer
RUN go build .

EXPOSE 8080
ENTRYPOINT ["./todo-consumer"]
FROM golang:latest

WORKDIR $GOPATH/src/yom535.coding.net/todo-select-provider
COPY . $GOPATH/src/yom535.coding.net/todo-select-provider
RUN go build .

EXPOSE 20880
ENTRYPOINT ["./todo-select-provider"]
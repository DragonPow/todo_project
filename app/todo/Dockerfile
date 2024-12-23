FROM golang:1.22.4 as builder

WORKDIR /

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux

COPY /third_party ./third_party

COPY /library ./library

COPY /go.mod /go.sum ./

RUN go mod download

COPY ./app/todo ./app/todo

RUN go build -o todo-web ./app/todo/cmd

FROM alpine:3.12

WORKDIR /app
COPY --from=builder /app/todo/config.yaml .
COPY --from=builder /todo-web .
EXPOSE 8080
CMD [ "/app/todo-web", "server" ]
FROM golang:latest

ADD . /go/src/app

WORKDIR /go/src/app

RUN go mod init
RUN go mod tidy

RUN env GOOS=linux GOARCH=amd64 go build -o build/main

ENV PORT=8080

CMD ["/go/src/app/build/main"]
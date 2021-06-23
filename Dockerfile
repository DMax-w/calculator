FROM golang:latest

ADD . /go/src/app

WORKDIR /go/src/app

RUN go get github.com/DMax-w/calculatorP
RUN go mod init
RUN env GOOS=linux GOARCH=amd64 go build -o build/main

CMD ["/go/src/app/build/main"]
FROM golang:1.19.1-alpine

RUN apk update && apk add git

WORKDIR /go/src

COPY ./src .

RUN go get github.com/gin-gonic/gin
RUN go get github.com/jszwec/csvutil

CMD [ "go", "run", "main.go" ]

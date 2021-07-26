# Builder
FROM golang:1.16.5-alpine3.14 as builder

RUN apk update && apk upgrade && \
    apk --update add git make

WORKDIR /app

COPY . .

RUN go build -o api.exe main.go

CMD /app/api.exe
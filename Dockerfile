FROM golang:alpine

WORKDIR /go/src/social-reports
ENV TZ=Europe/Moscow

ADD . .

# For change tz
RUN apk add --no-cache tzdata

EXPOSE 3000

CMD go build -o ./bin/server main.go &&\
    go build -o ./swag /go/pkg/mod/github.com/swaggo/swag@v1.7.0/cmd/swag/main.go &&\
    ./swag init &&\
    ./bin/server

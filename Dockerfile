FROM golang:1.16-alpine AS builder

RUN go version

COPY . /tgSnWeatherBot/
WORKDIR /tgSnWeatherBot/

RUN go mod download
RUN GOOS=linux go build -o ./.bin/bot /cmd/bot/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /tgSnWeatherBot/.bin/bot .

EXPOSE 80

CMD ["./bot"]
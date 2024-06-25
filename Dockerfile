FROM golang:1.20.3-alpine as golang-build

RUN apk update && apk upgrade
RUN apk add --no-cache git

WORKDIR /golang

COPY . .
RUN go mod download && go mod tidy
RUN go build -o ./main ./cmd/main/main.go

FROM alpine:3.20.1

WORKDIR /app
RUN apk update
RUN apk add --no-cache tzdata

EXPOSE 8080
ENV TZ=Asia/Makassar
RUN cp /usr/share/zoneinfo/$TZ /etc/localtime 


COPY --from=golang-build /golang/main /app/main

CMD ["./main"]
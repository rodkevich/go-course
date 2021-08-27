# build stage
FROM golang:1.16 AS builder
LABEL maintainer="ivan.rodkevich"

ENV GO111MODULE=auto
ENV CGO_ENABLED=0
ENV GOOS=linux

RUN go get -d -v github.com/google/uuid
#RUN dpkg-reconfigure ca-certificates

WORKDIR /go/src/github.com/rodkevich/go-course
COPY . .
WORKDIR /go/src/github.com/rodkevich/go-course/cmd/hw_weather_service/gateway

RUN go build -ldflags="-s -w" ./gateway.go

# server stage
FROM scratch
LABEL maintainer="ivan.rodkevich/gateway"

COPY --from=builder /go/src/github.com/rodkevich/go-course/cmd/hw_weather_service/gateway .
#COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 10000

CMD ["./gateway"]

# build stage
FROM golang:1.16 AS builder
LABEL maintainer="ivan.rodkevich"

ENV GO111MODULE=auto
ENV CGO_ENABLED=0
ENV GOOS=linux

RUN go get -d -v github.com/gin-gonic/gin
RUN go get -d -v github.com/elastic/go-elasticsearch

RUN #dpkg-reconfigure ca-certificates

WORKDIR /go/src/github.com/rodkevich/go-course
COPY . .
WORKDIR /go/src/github.com/rodkevich/go-course/cmd/hw_weather_service/weather

RUN go build -ldflags="-s -w" ./main.go

# server stage
FROM scratch
LABEL maintainer="ivan.rodkevich/weather"

COPY --from=builder /go/src/github.com/rodkevich/go-course/cmd/hw_weather_service/weather .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 9090

CMD ["./main"]

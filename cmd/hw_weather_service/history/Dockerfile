# build stage
FROM golang:1.16 AS builder
LABEL maintainer="ivan.rodkevich"

ENV GO111MODULE=auto
ENV CGO_ENABLED=0
ENV GOOS=linux

RUN go get -d -v github.com/gin-gonic/gin
RUN go get -d -v github.com/elastic/go-elasticsearch

WORKDIR /go/src/github.com/rodkevich/go-course
COPY . .
WORKDIR /go/src/github.com/rodkevich/go-course/cmd/hw_weather_service/history

RUN go build -ldflags="-s -w" ./history.go

# server stage
FROM scratch
LABEL maintainer="ivan.rodkevich/history"

COPY --from=builder /go/src/github.com/rodkevich/go-course/cmd/hw_weather_service/history .

EXPOSE 9091

CMD ["./history"]

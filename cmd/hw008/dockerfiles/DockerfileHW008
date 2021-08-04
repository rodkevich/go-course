# build stage
FROM golang:1.16 AS builder
LABEL maintainer="ivan.rodkevich"

ENV GO111MODULE=auto
ENV CGO_ENABLED=0
ENV GOOS=linux
#ENV GOBIN=/go/bin/

RUN go get -d -v github.com/gin-gonic/gin

WORKDIR /go/src/github.com/rodkevich/go-course
COPY . .
WORKDIR /go/src/github.com/rodkevich/go-course/cmd/hw008/app

RUN go build -ldflags="-s -w" ./main.go

# server stage
FROM scratch

COPY --from=builder /go/src/github.com/rodkevich/go-course/cmd/hw008/app .

EXPOSE 5051

CMD ["./main"]

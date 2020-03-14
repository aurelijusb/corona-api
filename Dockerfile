FROM golang:1.14 AS builder
WORKDIR /go/src/github.com/aurelijusb/corona-api
COPY cmd cmd
COPY LICENSE.md LICENSE.md
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/api cmd/api/main.go

FROM alpine:latest  
WORKDIR /root/
COPY --from=builder /go/src/github.com/aurelijusb/corona-api/bin/api .
COPY --from=builder /go/src/github.com/aurelijusb/corona-api/LICENSE.md .
ENV SERVER_HOST=0.0.0.0
ENV SERVER_PORT=80
CMD ["./api"]
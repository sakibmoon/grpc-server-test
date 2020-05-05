FROM golang:1.14.2 as builder

WORKDIR /go/grpc-server/src/

COPY . .

#RUN go get
RUN CGO_ENABLED=0 GOOS=linux go build  -o grpc-server -a -installsuffix cgo main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/grpc-server/src/grpc-server .

CMD ["./grpc-server"]

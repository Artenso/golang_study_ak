FROM golang:1.22-alpine AS builder

COPY . /task3.2.2.2/
WORKDIR /task3.2.2.2/

RUN go mod download
RUN go build -o server main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /task3.2.2.2/server .
COPY --from=builder /task3.2.2.2/.env .

EXPOSE 8080

CMD ["./server"]
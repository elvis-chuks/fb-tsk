# syntax=docker/dockerfile:1

FROM golang:1.18-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN go build -o ./mercurie ../main.go

## Deploy
FROM alpine:3.11.3

WORKDIR /

COPY --from=builder /app/mercurie .

EXPOSE 5001

CMD [ "/mercurie" ]
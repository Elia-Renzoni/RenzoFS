FROM golang:1.20 AS builder

WORKDIR /app

RUN go mod download
RUN go build -o RenzoFS-API-Gateway

#TODO
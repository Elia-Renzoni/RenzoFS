# docker image for the authentication microservice

FROM golang:1.20

# destination of the sources in the file system
WORKDIR /DFSAuthMicro

COPY go.mod go.sum ./ 

RUN go mod download

# copy all the go files in /DFS
COPY *.go ./

# compile all the go files
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

# TCP port
EXPOSE 8082

# run the binary
CMD ["/docker-gs-ping"]
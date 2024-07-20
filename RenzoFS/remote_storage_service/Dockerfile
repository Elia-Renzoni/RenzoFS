# docker image for the remote storage microservice

FROM golang:1.20

# destination of the sources in the file system
WORKDIR /RenzoFS

COPY go.mod go.sum ./ 

COPY /local_file_system ./

RUN go mod download

# copy all the go files in /DFS
COPY *.go ./

# compile all the go files
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

# TCP port
EXPOSE 8080

# run the binary
CMD ["/docker-gs-ping"]
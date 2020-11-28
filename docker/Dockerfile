FROM golang
WORKDIR /app
COPY . .
RUN go mod download
#complieDaemon for live reload golang
RUN go get github.com/githubnemo/CompileDaemon
ENTRYPOINT CompileDaemon --build="go build server.go" --command=./server
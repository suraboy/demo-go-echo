FROM golang
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

RUN go get github.com/githubnemo/CompileDaemon

#CMD ["go","run","server.go"]
ENTRYPOINT CompileDaemon --build="go build server.go" --command=./server
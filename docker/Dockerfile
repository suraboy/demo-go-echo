FROM golang
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GO111MODULE=on GOSUMDB=off go build -mod=vendor -mod=mod -a -o app .

RUN go get github.com/githubnemo/CompileDaemon

#CMD ["go","run","server.go"]
ENTRYPOINT CompileDaemon --build="go build server.go" --command=./server
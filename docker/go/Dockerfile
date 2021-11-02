FROM golang:latest

WORKDIR /app

# чтоб эти файлы взять из родительской директории
# настроили котекст в docker-compose.yml
# иначе при запуске файлы не видны и не запускался "go mod download"
COPY go.sum .
COPY go.mod .

ENV GO111MODULE=on

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon -build="go build server.go" -command=./server
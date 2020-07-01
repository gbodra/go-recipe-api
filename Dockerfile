  
FROM golang:1.14

RUN mkdir /app
ADD . /app/
WORKDIR /app

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon -log-prefix=false -build="make" -command="./go-recipe-api"
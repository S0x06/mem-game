#!/bin/sh

echo "### Build project ..."
go get github.com/emirpasic/gods/lists/arraylist
CGO_ENABLED=0 go build -a -installsuffix cgo -o em-game-server main.go

echo "### Build docker image ..."
docker build -t em-game-server:latest .

echo "### Package docker image ..."
docker save em-game-server:latest | gzip > em-game-server.tgz
echo "### Package docker image finish."
#!/bin/sh

echo "### Build project ..."
go get github.com/emirpasic/gods/lists/arraylist
go build --tags "libsqlite3 linux" -o mem-game-server main.go

echo "### Build docker image ..."
docker build -t mem-game-server:latest .

echo "### Package docker image ..."
docker save mem-game-server:latest | gzip > mem-game-server.tgz
echo "### Package docker image finish."

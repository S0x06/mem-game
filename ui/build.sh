#!/bin/sh

echo "### Build docker image ..."
docker build -t mem-game-ui:latest .

echo "### Package docker image ..."
docker save mem-game-ui:latest | gzip > mem-game-ui.tgz
echo "### Package docker image finish."


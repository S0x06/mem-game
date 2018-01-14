#!/bin/sh

echo "### Build docker image ..."
docker build -t em-game-ui:latest .

echo "### Package docker image ..."
docker save em-game-ui:latest | gzip > em-game-ui.tgz
echo "### Package docker image finish."